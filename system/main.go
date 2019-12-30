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
	c.Set("STARTTIME", time.Now().String(), 0)

}
