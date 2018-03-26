package cli

import "errors"
import "strconv"
import "strings"
import "github.com/go-redis/redis"

type Cli struct {
	client redis.Client
	cursor uint64
	start  bool
}

func (c *Cli) Next() ([]string, error) {
	var keys []string
	var cur uint64
	var err error
	if c.start {
		if c.cursor == 0 {
			c.start = false
			return keys, errors.New("no more keys")
		} else {
			keys, cur, err = c.client.Scan(c.cursor, "", 10).Result()
			if err != nil {
				c.cursor = 0
				c.start = false
				return []string{}, errors.New("scan key err")
			} else {
				c.cursor = cur
				return keys, nil
			}
		}
	} else {
		keys, cur, err = c.client.Scan(0, "", 10).Result()
		if err != nil {
			c.cursor = 0
			c.start = false
			return []string{}, errors.New("scan key err")
		} else {
			c.cursor = cur
			c.start = true
			return keys, nil
		}

	}
}

func (c *Cli) GetLength(key string) (int, error) {
	ret, err := c.client.DebugObject(key).Result()
	if err != nil {
		return 0, err
	}
	decode := func(str string) (int, error) {
		strs := strings.Fields(str)
		for _, str := range strs {
			if strings.HasPrefix(str, "serializedlength") {
				cobIdx := strings.Index(str, ":")
				num, err := strconv.Atoi(str[cobIdx+1:])
				if err != nil {
					return 0, err
				}
				return num, nil
			}
		}
		return 0, errors.New("decode error")
	}
	return decode(ret)
}

func CliNew(addr string, pswd string) Cli {
	c := Cli{}
	c.client = *redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pswd,
		DB:       0,
	})
	c.cursor = 0
	c.start = false
	return c
}
