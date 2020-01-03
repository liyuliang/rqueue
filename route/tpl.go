package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/system"
	"github.com/liyuliang/utils/format"
	"github.com/pkg/errors"
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"encoding/base64"
)

func tpl(c *gin.Context) {

	tplDir := system.Config()[system.SystemTplDir]

	if tplDir == "" {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   errors.New("wrong tpl dir ").Error(),
			"success": "false",
		}))
		return
	}

	ext := ".toml"

	var files []string

	err := filepath.Walk(tplDir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		c.JSON(200, format.ToMap(map[string]string{
			"error":   err.Error(),
			"success": "false",
		}))
	}

	m := format.Map()

	for _, file := range files {

		data, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}

		k := strings.Replace(file, tplDir, "", -1)
		k = strings.Replace(k, ext, "", -1)
		k = strings.Replace(k, "/", "", -1)


		//m[k] = string(data)
		m[k] = base64.StdEncoding.EncodeToString(data)
	}

	m["error"] = ""
	m["success"] = "true"
	c.JSON(200, m)
}
