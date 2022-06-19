package main

/*
cs@debian:~/go/src/demo$ go mod init demo
go: creating new go.mod: module demo
go: to add module requirements and sums:
	go mod tidy
*/

import (
	"demo/com/cs/learn/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.ReleaseMode

}

func newFunction() {
	gin.SetMode(gin.ReleaseMode)
	rt := gin.Default()
	router.InitRouter(rt)
	rt.Run()
}
