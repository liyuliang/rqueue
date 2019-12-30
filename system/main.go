package system

import (
	"time"
	"github.com/liyuliang/utils/format"
)

func Init(data format.MapData) {

	c = data

	initRedis()
}

func initRedis() {

	c := Redis()
	c.Set("START_TIME", time.Now().String(), 0)

}
