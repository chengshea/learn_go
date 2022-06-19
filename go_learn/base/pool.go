package bs

import (
	"fmt"
	"sync"
)

type student struct {
	Name string
	Age  int
}

//存储临时对象
var stuP = &sync.Pool{
	New: func() interface{} {
		return new(student)
	},
}

func get(n string, a int) *student {
	stu := stuP.Get().(*student)
	stu.Name = n + " 业务"
	stu.Age = a + 2
	return stu
}

func clear(stu *student) {
	stu.Name = ""
	stu.Age = 0
	stuP.Put(stu)
}

func Pool() {
	stu := get("zs", 11)
	fmt.Println(*stu)
	defer clear(stu) // 释放
}
