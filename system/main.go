package system

import "time"

func Init(redisUri string) {

	_config["redisUri"] = redisUri

	initRedis()
}

func initRedis() {

	c := Client()
	c.Set("STARTTIME", time.Now().String(), 0)

}
