/*
安装Go语言开发工具包 ctrl+Shift+P
输入框中输入go:install 会自动搜索相关命令，选择Go:Install/Update Tools
设置代理
$ go env -w GOPROXY=https://goproxy.cn,direct
cd ./study
go mod init go_learn


$ go env -w GO111MODULE=auto
$ go env | grep MODULE
GO111MODULE="auto"



*/

package main

import (
	"fmt"
	bs "go_learn/base"
	t1 "go_learn/test"
	tr "go_learn/training"
	"os"
)

func main() {
	// mf()
	//	gaojf()
	//bs.SyncMap()
	//bs.Pool()
	pwd, _ := os.Getwd()
	fmt.Println("main", pwd)
	//training()

	//test0()
}

func test0() {
	ia := make([]int, 256)
	ia[0] = 0
	ia[3] = 's'
	ia[4] = 'a'
	fmt.Println(ia[3], ia)
}

func training() {
	//	tr.Palter()
	//tr.Plongest()
	tr.Plong()
}

func gaojf() {
	//bs.Struct1()

	//bs.Js()
	//	bs.AutoGen()

	//	bs.Chanstr()//chan

	//	bs.UseDefer() //defer 类似java的finally
	bs.InterfaceImpl()
	//bs.Search()
	//bs.UseCom()
}

func basef() {
	//bs.Arr1()
	//	bs.Map1()
	//bs.Map2()
}

func mf() {
	fmt.Println("main调用包test fun1")

	t1.Tfun1() //两个返回值
	//t1.Zz() //指针
}
