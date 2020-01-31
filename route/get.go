package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/rqueue/system"
	"strings"
)

func get(c *gin.Context) {

	queue, err := getPostParam(c, "queue")

	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error() + ":" + system.QueueInGetApi,
			"success": "false",
		}))
		return
	}

	num, _ := c.GetPostForm("n")

	n := format.StrToInt(num)
	if n < 1 {
		n = format.StrToInt(system.Config()["popNum"])
	}

	var data []string
	client := system.Redis()
	for i := 0; i < n; i++ {
		v := client.LPop(queue).Val()
		if v == "" {
			break
		}

		data = append(data, v)
		k := strings.Replace(queue, system.RedisQueuePrefix, system.RedisQueueTotalPrefix, -1)
		client.Decr(k)
	}
	c.JSON(200, data)
}
