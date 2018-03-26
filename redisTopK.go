package main

import "fmt"
import "flag"

import "redisTopK/cli"
import "redisTopK/topk"

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

var (
	host string
	pswd string
	k    int
)

func main() {
	// parse args
	flag.StringVar(&host, "host", "127.0.0.1:6379", "redis host to connect")
	flag.StringVar(&pswd, "pswd", "", "redis host password to connect")
	flag.IntVar(&k, "k", 10, "top k")
	flag.Parse()

	tk := topk.TopKNew(k)
	client := cli.CliNew(host, pswd)
	var keys []string
	var err error
	keys, err = client.Next()
	for err == nil {
		for _, key := range keys {
			length, err := client.GetLength(key)
			if err == nil {
				// insert into tk table
				tk.Insert(Info{key, length})
			}
		}
		keys, err = client.Next()
	}
	fmt.Printf("redis %v top %v value:\n", host, k)
	tk.Print()
}
