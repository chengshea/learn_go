package bs

import (
	"fmt"
	"sync"
)

var (
	cache = &sync.Pool{
		New: func() interface{} {
			return &option{sex: 0}
		},
	}
)

func (o *option) reset() {
	o.sex = 0
	o.age = 0
	o.height = 0
	o.weight = 0
	o.like = ""
}

type Option func(*option)
type option struct {
	sex    int
	age    int
	height int
	weight int
	like   string
}

func getOption() *option {
	return cache.Get().(*option)
}
func clearOption(o *option) {
	o.reset()
	cache.Put(o)
}

func setSex(s int) Option {
	return func(o *option) {
		o.sex = s
	}
}

func setAge(a int) Option {
	return func(o *option) {
		o.age = a
	}
}

func setHight(h int) Option {
	return func(o *option) {
		o.height = h
	}
}
func setWight(w int) Option {
	return func(o *option) {
		o.weight = w
	}
}

func setLike(l string) Option {
	return func(o *option) {
		o.like = l
	}
}

func find(str string, opt ...Option) (string, error) {
	fs := fmt.Sprintf("从%s搜索\n", str)
	o := getOption()
	//defer func() { clearOption(o) }()//清除释放
	defer clearOption(o)
	fmt.Println("get:", o)
	//赋值
	for _, val := range opt {
		val(o)
	}
	fmt.Println("val:", o)

	if o.sex == 0 {
		sex := "性别:女"
		fs += fmt.Sprintf("%s\n", sex)
	} else {
		sex := "性别:男"
		fs += fmt.Sprintf("%s\n", sex)
	}

	fs = juge(o.age, "年龄", fs)
	fs = juge(o.weight, "体重", fs)
	fs = juge(o.height, "身高", fs)
	fs = juge(o.like, "爱好", fs)
	return fs, nil
}

//断言的基础使用x.(T)
//i ...interface{}
func juge(i interface{}, s string, fs string) string {
	var val string
	switch i.(type) {
	case int:
		val = fmt.Sprintf("%s:%d", s, i)
	case string:
		val = fmt.Sprintf("%s:%s", s, i)
	default:
		fmt.Println("类型异常")
	}
	fs += fmt.Sprintf("%s\n", val)
	return fs
}

func Search() {
	fs, err := find(
		"附近人",
		setSex(0),
		setAge(12),
		setWight(123),
		setHight(111),
		setLike("live"))

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(fs)
}
