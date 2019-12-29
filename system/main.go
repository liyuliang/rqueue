package system

import "time"

func Init(redisUri string) {

	c := Conn(redisUri)
	c.Set("name", "liang", time.Second * 600)
}
