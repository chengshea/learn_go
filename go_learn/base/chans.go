package bs

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	//	"go_learn/tool"
)

func Chanstr() {

	ch2 := make(chan string, 8)
	// ch3 := make(<-chan string)
	// ch4 := make(chan<- string)

	go producer(ch2)

	go customer(ch2)
	time.Sleep(2 * time.Second)
}

func bfstr(bf bytes.Buffer, str string, i string) string {
	bf.WriteString(str)
	bf.WriteString(i)
	return bf.String()
}

func producer(ch2 chan string) {
	fmt.Println("producer start")
	var bf bytes.Buffer
	for i := 0; i < 7; i++ {

		ch2 <- bfstr(bf, "ch2-", strconv.Itoa(i))
	}
	fmt.Println("producer end")
}

func customer(ch chan string) {
	bf := bytes.NewBufferString("")
	for {
		msg := <-ch
		fmt.Println(msg)
		bf.WriteString(msg)
		bf.WriteString(",")
		if len(ch) == 0 {
			//去掉最后一个字符
			fmt.Printf("bf TrimRight: %v\n", strings.TrimRight(bf.String(), ","))
			fmt.Printf("bf len-1: %v\n", bf.String()[:len(bf.String())-1])
			break
		}
	}
}
