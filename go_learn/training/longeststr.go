package training

import (
	"fmt"
	"strings"
	"time"
)

var m map[string]int
var sb strings.Builder
var res string

func Plongest() {
	stt := time.Now()
	str := "abcdbcdeafghcedabcdefg"

	handleStr(str)
	defer func() {
		cost := time.Since(stt)
		fmt.Printf("len:%d res:%s time:%s\n", len(res), res, cost)
	}()
}

func handleStr(str string) {
	b := []byte(str)
	m = make(map[string]int)
	for i := range b {
		t := string(b[i])
		fmt.Println("st:", t)
		if i > 0 {
			if juge(t) {
				fmt.Println("end:", m, sb.String())
				continue
			}

		}
		sb.WriteString(t)
		m[t] = i
		tempsto(sb.String())
		fmt.Println("end:", m, sb.String())
	}
}

/*
判断是否存在
*/
func juge(str string) bool {
	_, flag := m[str]
	if flag {
		temp := sb.String()
		x := strings.Split(temp, str)
		fmt.Println("falg..", flag, str, temp, x)
		l := len(x)
		if l > 1 {
			sb = strings.Builder{} //重复清空
			sb.WriteString(x[1])
			del(x[0])

		}
		sb.WriteString(str)
		return flag
	}
	return flag
}

/*
删除map里重复字符
*/
func del(str string) {
	b := []byte(str)
	for i := range b {
		delete(m, string(b[i]))
	}
}

/*
存储最长字串
*/
func tempsto(str string) {
	fmt.Println("存储;", res, str)
	if len(str) > len(res) {
		res = str
	}

}
