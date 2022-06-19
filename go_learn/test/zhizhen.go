package t1

import "fmt"

func zz(i *int) {
	*i = 8
	fmt.Println("*i", *i)
}
func z(i int) {
	var a *int
	a = &i
	fmt.Println("&i,*a:", a, *a)
}
func Zz() {
	val := 2
	z(val)
	zz(&val)
	fmt.Println(val)
}
