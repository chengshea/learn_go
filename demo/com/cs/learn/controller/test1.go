package controller

import (
	"demo/com/cs/learn/tool"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
http://localhost:8080/v1/q?name=s
*/
func Tget1(c *gin.Context) {
	name := c.Query("name")
	req := c.Request.URL.Query()
	c.JSON(200,
		gin.H{
			"v1":   "Test1",
			"name": name,
		})
	log.Println("Test1", req)
}

func Tget2(c *gin.Context) {
	name := c.Query("name")
	if !tool.IsStrNull(name) {
		client := tool.OpenConRedis()
		val, err := client.Get(ctx, name).Result()
		if !tool.IsExist(err) {
			log.Fatalln("key does not exist")
		} else if err != nil {
			log.Fatalln(val, err)
		}
		name = val
		//	defer client.Close()
	}

	c.JSON(http.StatusOK,
		gin.H{
			"v1":   "Test1",
			"name": name,
		})
}

/*
// 获取 Get 参数
name := c.Query("name")
price := c.DefaultQuery("price", "100")

// 获取 Get 所有参数
ReqGet = c.Request.URL.Query()

*/
