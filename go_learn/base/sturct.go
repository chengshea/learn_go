package bs

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}

func Struct1() {
	var s1 Student
	s1.Name = "tol"
	s1.Age = 11
	fmt.Println("s1:", s1)
	s2 := Student{Name: "cs", Age: 21}
	fmt.Println("s2:", s2)
	// 匿名
	s3 := struct {
		Name string
		Age  int
	}{Name: "lisi", Age: 33}
	fmt.Println("s3:", s3)

	p := unsafe.Pointer(&s3)
	name := (*string)(unsafe.Pointer(p))
	age := unsafe.Offsetof(s3.Age)
	fmt.Println("unsafe:", *name, age)

}

type Blog1 struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     string `mapstructure:"uid"`
	State   string `mapstructure:"state"`
}

type Blog2 struct {
	BlogId  string `mapstructure:"blogId"`
	Title   string `mapstructrue:"title"`
	Content string `mapstructure:"content"`
	Uid     int32  `mapstructure:"uid"`
	State   int32  `mapstructure:"state"`
}

type Event struct {
	Type     string              `json:"type"`
	Database string              `json:"database"`
	Table    string              `json:"table"`
	Data     []map[string]string `json:"data"`
}

type AutoGenBlog struct {
	Type     string `json:"type"`
	Database string `json:"database"`
	Table    string `json:"table"`
	Data     []struct {
		BlogID  string `json:"blogId"`
		Title   string `json:"title"`
		Content string `json:"content"`
		UID     string `json:"uid"`
		State   string `json:"state"`
	} `json:"data"`
}
