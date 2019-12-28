package route

import (
	"github.com/gin-gonic/gin"
)

func Start(port string) {

	r := gin.Default()
	r.GET("/add", add)


	r.NoRoute(method404)
	r.Run(":" + port)
}
