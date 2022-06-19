package t1

import (
	"fmt"
)

var gloal string = " 全局"

func Tfun2(a ...string) (string, string) {
	temp := " 局部"
	nf(a)
	return a[0] + gloal, a[1] + temp
}

func nf(a []string) {
	for i, v := range a {
		if a[i] == "dd" {
			fmt.Println(a, v)
		}
	}
}
