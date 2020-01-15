package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/utils/format"
	"github.com/pkg/errors"
	"github.com/liyuliang/rqueue/system"
	"log"
)

func getAuthData(c *gin.Context) format.MapData {
	data := format.Map()

	data["ip"] = c.ClientIP()
	data["host"] = c.PostForm("host")
	data["system"] = c.PostForm("system")
	data["core"] = c.PostForm("core")
	data["load"] = c.PostForm("load")
	data["memory"] = c.PostForm("memory")
	data["disk"] = c.PostForm("disk")
	return data
}

func auth(c *gin.Context) {

	data := getAuthData(c)

	err := checkAuthData(data)

	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error": "Auth param wrong",
			"uuid":  "",
		}))
		return
	}

	id, err := system.Uid()
	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error": err.Error(),
			"uuid":  "",
		}))
		return
	}

	data["uuid"] = format.Int64ToStr(id)

	client := system.Redis()
	r := client.HSet(system.SpiderClientSet, c.ClientIP(), data.String())
	if r.Err() != nil {
		log.Print(r.Err().Error())
	}

	c.JSON(200, format.ToMap(map[string]string{
		"error": "",
		"uuid":  format.Int64ToStr(id),
	}))
}

func checkAuthData(data format.MapData) (err error) {
	for key, val := range data {
		if val == "" {
			err = errors.New("Key :" + key + " in auth is required")
			break
		}
	}
	return err
}
