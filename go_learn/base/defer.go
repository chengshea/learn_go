package bs

import "fmt"

func UseDefer() {
	a := 1
	b := 2

	fmt.Println("start a b:", a, b)

	defer func() { fmt.Println("1闭包defer a+b:", a+b) }()

	defer func(a int, b int) {
		fmt.Println("2defer(a,b) 前 a+b:", a+b)
	}(a, b)
	/*
		def(a,b) 后 a+b: 3
		def 闭包 a+b: 3
	*/
	val := def(a, b)
	fmt.Println("func return val:", val)

	a = 2
	fmt.Println("中间a:", a)
	/*
		return不影响,闭包随函数影响
		def2(a,b) 后 a+b: 4
		def2 闭包 a+b: 4
	*/
	def2(a, b)
	defer func(a int, b int) {
		fmt.Println("3defer(a,b) 后 a+b:", a+b)
	}(a, b)
	defer fmt.Println("4闭包defer a+b:", a+b)
	fmt.Println("end a b:", a, b)

	/*

		start a b: 1 2
		中间a: 2
		end a b: 2 2
		4defer fmt a+b: 4
		3defer(a,b) 后 a+b: 4
		2defer(a,b) 前 a+b: 3
		1defer fmt a+b: 4

		闭包 引用传递
		传参数  值的复制,不影响实参
	*/
}

func def(a int, b int) int {
	defer fmt.Println("def 闭包 a+b:", a+b)
	defer func(a int, b int) {
		fmt.Println("def(a,b) 后 a+b:", a+b)
	}(a, b)
	return a + b
}

func def2(a int, b int) {
	defer fmt.Println("def2 闭包 a+b:", a+b)
	defer func(a int, b int) {
		fmt.Println("def2(a,b) 后 a+b:", a+b)
	}(a, b)
}
