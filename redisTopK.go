package main

import "fmt"
import "flag"
import "sync"
import "time"

import "redisTopK/cli"
import "redisTopK/topk"

var (
	host string
	pswd string
	k    int
)

var wg sync.WaitGroup

type Info struct {
	name string
	size int
}

func (i Info) Val() int {
	return i.size
}

func (i Info) Str() string {
	str := fmt.Sprintf("<%v:%v>", i.name, i.size)
	return str
}

// get keys from db
func getKeys(db int, c chan Info) {
	wg.Add(1)
	client := cli.CliNew(host, pswd, db)
	var keys []string
	var err error
	keys, err = client.Next()
	for err == nil {
		for _, key := range keys {
			length, err := client.GetLength(key)
			if err == nil {
				c <- Info{key, length}
			}
		}
		keys, err = client.Next()
	}
	c <- Info{"", -1}
}

// update statistics
func stats(t *topk.TopK, c chan Info) {
	for {
		i := <-c
		if i.size != -1 {
			t.Insert(i)
		} else {
			wg.Done()
		}
	}
}

func main() {
	// parse args
	flag.StringVar(&host, "host", "127.0.0.1:6379", "redis host to connect")
	flag.StringVar(&pswd, "pswd", "", "redis host password to connect")
	flag.IntVar(&k, "k", 10, "top k")
	flag.Parse()

	client := cli.CliNew(host, pswd, 0)
	dbs := client.GetDatabases()
	if dbs == 0 {
		fmt.Printf("redis has no db")
		return
	}

	c := make(chan Info, 100)
	tk := topk.TopKNew(k)
	go stats(tk, c)
	for j := 0; j < dbs; j++ {
		go getKeys(j, c)
	}

	time.Sleep(100 * time.Millisecond)
	wg.Wait()
	tk.Print()
	close(c)
}
