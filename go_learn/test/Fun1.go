package t1

import (
	"fmt"
)

func Tfun1() {
	a, b := Tfun2("a is", "b is", "dd")
	fmt.Println(a, b)
	str := []string{"dd", "bd", "ddd"}
	nf(str)
}
