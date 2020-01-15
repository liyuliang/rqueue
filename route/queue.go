package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
	"strings"
	"github.com/liyuliang/utils/format"
)

func queue(c *gin.Context) {

	client := system.Redis()

	data := format.Map()
	for _, v := range client.Keys("*").Val() {
		if strings.Contains(v, system.RedisQueuePrefix) {

			k := strings.Replace(v, system.RedisQueuePrefix, system.RedisQueueTotalPrefix, -1)
			data[v] = client.Get(k).Val()
		}
	}
	c.JSON(200, data)
}
