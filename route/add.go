package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/request"
	//"github.com/gorilla/mux"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"net/url"
	"github.com/pkg/errors"
)

const URL = "url"
const CATEGORY = "type"

func add(c *gin.Context) {

	data := request.Data(c.Request)

	err := checkJob(data)
	if err != nil {

	}

	key := genQueueKey(data)
	client := system.Client()
	r := client.RPush(key, format.ToMapData(data).ToString())
	if r.Err() != nil {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   r.Err().Error(),
			"success": "false",
		}))
	} else {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   "",
			"success": "true",
		}))
	}
	//routeParams := mux.Vars(c.Request)
}

func genQueueKey(data map[string]string) string {
	url := data[URL]
	category := data[CATEGORY]
	return category + getKeyFromUrl(url)
}

func checkJob(data map[string]string) error {

	url := data[URL]
	category := data[CATEGORY]

	key := getKeyFromUrl(url)

	if url == "" {
		return errors.New("url is required")
	}
	if category == "" {
		return errors.New("type is required")
	}
	if key == "" {
		return errors.New("It's not avail url")
	}
	return nil
}

func getKeyFromUrl(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	if u.Scheme == "" {
		return ""
	}
	if u.Host == "" {
		return ""
	}
	key := regex.Get(uri, `\/\/[^\.]+\.([^\.]+)\.`)
	return key
}
