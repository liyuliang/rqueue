package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
)

func uuid(c *gin.Context) {

	id, err := system.Uid()
	if err != nil {

		c.JSON(200, format.ToMap(map[string]string{
			"error": err.Error(),
			"uuid":  "",
		}))
	} else {

		c.JSON(200, format.ToMap(map[string]string{
			"error": "",
			"uuid":  format.Int64ToStr(id),
		}))
	}
}
