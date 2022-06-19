package controller

import (
	"context"
	"demo/com/cs/learn/tool"
	"log"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

/*
http://localhost:8080/v2/q
*/
func Tpost1(c *gin.Context) {
	name := c.PostForm("name")
	rq := c.Request.PostForm
	c.JSON(200,
		gin.H{
			"v2":   "Test2",
			"name": name,
		})
	log.Fatalln("Test2", rq)
}

func Tpost2(c *gin.Context) {
	name := c.PostForm("name")
	if !tool.IsStrNull(name) {
		client := tool.OpenConRedis()
		err := client.Set(ctx, name, "1234", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	c.JSON(200,
		gin.H{
			"v2":   "Tpost2",
			"name": name,
		})

}

/*
// 获取 Post 参数
name := c.PostForm("name")
price := c.DefaultPostForm("price", "100")
//获取 Post 所有参数
ReqPost = c.Request.PostForm

*/
