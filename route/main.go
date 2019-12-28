package route

import (
	"github.com/gin-gonic/gin"
)

func Start(port string) {

	r := gin.Default()

	r.NoRoute(method404)
	r.Run(":" + port)
}
