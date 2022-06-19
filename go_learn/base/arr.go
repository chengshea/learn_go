/**

首字母大写则是公开的名称，导入包后可以直接引用，小写则为私有
构造函数一般会成为NewXXX，如果是包中唯一类型，则也可以只叫做New

*/
package bs

import "fmt"

func Arr1() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [5]int{0: 3, 1: 5, 4: 6}
	//arr1:[1 2 3 4 5] arr2:[1 2 3 4 5] arr3:[3 5 0 0 6]
	fmt.Printf("arr1:%d arr2:%d arr3:%d\n", arr1, arr2, arr3)
	modifyArr(arr1)

	modifyArr2(&arr3) //modify arr1:[1 2 3 4 5] arr3:[20 5 0 0 6]
	fmt.Printf("modify  arr1:%d &arr3:%d\n", arr1, arr3)

	arr4 := make([]int, 3, 6) //make len=3 cap=6
	fmt.Printf("make len=%d cap=%d\n", len(arr4), cap(arr4))
	//len=5 cap=5
	printArr1("arr2", arr2)
	//len=2 cap=5
	printArr("arr2[0:2]", arr2[0:2])
	printArr("arr2[0:2:3]", arr2[0:2:3])

	arr5 := []int{1, 2, 3, 4, 5}
	arr := append(append(arr5, 6), 7)
	fmt.Printf("append %d\n", arr)
	fmt.Println("2: 左边开始截取   :len-2 右边开始截取")
	fmt.Printf("del {[1 2 3 4 5 6 7} [2:]=%d [:len(arr)-2]=%d\n", arr[2:], arr[:len(arr)-2])
}

func modifyArr(a [5]int) {
	a[0] = 20
}

func modifyArr2(a *[5]int) {
	a[0] = 20
}

func printArr1(str string, a [5]int) {
	fmt.Printf("val:%s  len:%d   cap:%d\n", str, len(a), cap(a))
}

func printArr(str string, a []int) {
	fmt.Printf("val:%s  len:%d   cap:%d    arr:%d\n", str, len(a), cap(a), a)
}
