package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
)

func keys(c *gin.Context) {

	client := system.Redis()

	var data []string
	for _, v := range client.Keys("*").Val() {
		if v != "" {
			data = append(data, v)
		}
	}
	c.JSON(200, data)
}
