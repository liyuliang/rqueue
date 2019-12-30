package route

import (
	"github.com/gin-gonic/gin"
)

func Start(port string) {

	r := gin.Default()
	r.GET("/keys", keys)
	r.GET("/queue", queue)
	r.GET("/uuid", uuid)
	r.POST("/auth", auth)
	r.POST("/add", add)
	r.POST("/get", get)

	r.NoRoute(method404)
	r.Run(":" + port)
}
