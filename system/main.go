package system

import (
	"time"
	"github.com/liyuliang/utils/format"
)

func Init(data format.MapData) {

	_config = data

	initRedis()
}

func initRedis() {

	c := Client()
	c.Set("START_TIME", time.Now().String(), 0)

}
