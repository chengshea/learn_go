package router

import (
	ct "demo/com/cs/learn/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(rt *gin.Engine) {

	v1 := rt.Group("/v1")
	{
		v1.GET("/g1", ct.Tget1)
		v1.GET("/g2", ct.Tget2)
	}
	v2 := rt.Group("/v2")
	{
		v2.POST("/p1", ct.Tpost1)
		v2.POST("/p2", ct.Tpost2)
	}
}
