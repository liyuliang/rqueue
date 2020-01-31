package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/utils/regex"
	"net/url"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func add(c *gin.Context) {

	//token
	//type
	//[]url

	token, err := getPostParam(c, system.TokenInAddApi)
	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error() + ":" + system.TokenInAddApi,
			"success": "false",
		}))
		return
	}

	token = system.RedisTokenExpiredPrefix + token
	err = checkUnique(token)

	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error(),
			"success": "false",
		}))
		return
	}

	cat, err := getPostParam(c, system.CategoryInAddApi)
	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error() + ":" + system.CategoryInAddApi,
			"success": "false",
		}))
		return
	}

	urls, err := getPostParams(c, system.UrlsInAddApi)
	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error() + ":" + system.UrlsInAddApi,
			"success": "false",
		}))
		return
	}

	client := system.Redis()
	for _, u := range urls {
		queue := genQueueName(cat, u)
		total := genTotalQueueName(cat, u)

		client.RPush(queue, u)

		//队列集-总数
		client.Incr(total)
	}

	r := client.Set(token, 120, 120*time.Second)

	c.JSON(200, format.ToMap(map[string]string{
		"error":   format.ErrorToStr(r.Err()),
		"success": "true",
	}))
}

func checkUnique(k string) error {

	redis := system.Redis()
	r := redis.Get(k)

	if r.Val() != "" && system.Config()[system.SystemIsDebug] == "false" {
		return errors.New("token exist, this is repeat request")
	}
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

func genQueueType(category, uri string) string {
	// same as tpl dir config file name
	return getKeyFromUrl(uri) + "_" + category
}

func genQueueName(category, uri string) string {
	return system.RedisQueuePrefix + genQueueType(category, uri)
}

func genTotalQueueName(category, uri string) string {
	return system.RedisQueueTotalPrefix + genQueueType(category, uri)
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
