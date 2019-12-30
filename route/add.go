package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"net/url"
	"github.com/pkg/errors"
	"strings"
)

const uuid = "uuid"
const uri = "url"
const uris = "[]url"
const category = "type"

func add(c *gin.Context) {

	id, err := getPostParam(c, uuid)
	if err != nil {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   err.Error() + ":" + uuid,
			"success": "false",
		}))
		return
	}

	err = checkUUid(id)

	if err != nil {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   err.Error(),
			"success": "false",
		}))
		return
	}

	cat, err := getPostParam(c, category)
	if err != nil {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   err.Error() + ":" + category,
			"success": "false",
		}))
		return
	}

	urls, err := getPostParams(c, uris)
	if err != nil {
		c.JSON(200, format.ToMapData(map[string]string{
			"error":   err.Error() + ":" + uris,
			"success": "false",
		}))
		return
	}

	client := system.Client()
	for _, u := range urls {
		key := genQueueKey(cat, u)

		client.RPush(key, u)
		//队列集-总数
		// client.HSet()
	}


	c.JSON(200, format.ToMapData(map[string]string{
		"error":   "",
		"success": "true",
	}))
}

func checkUUid(uuid string) error {
	//TODO 24小时 uuid, 防止重播攻击
	return nil
}

func getPostParams(c *gin.Context, key string) ([]string, error) {
	v := c.PostFormArray(key)
	if len(v) == 0 {
		v2 := c.PostForm(strings.Replace(key, "[]", "", -1))
		if v2 != "" {
			v = []string{v2}
		}
	}
	return v, checkUris(v)
}

func getPostParam(c *gin.Context, key string) (string, error) {
	v := c.PostForm(key)
	return v, checkEmpty(v)
}

func genQueueKey(category, uri string) string {
	return category + "_" + getKeyFromUrl(uri)
}

func checkEmpty(v string) error {
	if v == "" {
		return errors.New("missing required param")
	}
	return nil
}

func checkUri(uri string) error {

	if uri == "" {
		return errors.New("url is addRequired")
	}

	key := getKeyFromUrl(uri)

	if key == "" {
		return errors.New("It's not avail url")
	}
	return nil
}

func checkUris(uris []string) (err error) {
	for _, uri := range uris {

		err = checkUri(uri)
		if err != nil {
			break
		}
	}
	return err
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
