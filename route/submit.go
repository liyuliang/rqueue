package route

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"github.com/liyuliang/rqueue/system"
)

func submit(c *gin.Context) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	str := buf.String()

	k := system.RedisQueueStorage
	client := system.Redis()

	client.RPush(k, str)
	client.Incr(system.RedisQueueTotalPrefix + k)

	c.String(200, str)
}
