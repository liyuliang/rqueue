package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyuliang/rqueue/request"
	//"github.com/gorilla/mux"
)

func add(c *gin.Context) {

	request.Data(c.Request)
	//routeParams := mux.Vars(c.Request)

}
