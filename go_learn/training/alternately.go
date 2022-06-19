package training

import (
	"fmt"
	"sync"
)

var t chan int

func Palter() {
	str := "hello world"
	var wg sync.WaitGroup
	t = make(chan int)
	wg.Add(2)
	go alternately1(str, &wg)
	go alternately2(str, &wg)
	wg.Wait()
}

func alternately1(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	b := []byte(str)
	for i := range b {
		fmt.Printf("t:%d -", t)
		t <- 1
		if i%2 == 0 {
			fmt.Println("alternately1:", string(b[i]))
		}
	}
}

func alternately2(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	b := []byte(str)
	for i := range b {
		fmt.Printf("2t:%d -", t)
		<-t
		if i%2 == 1 {
			fmt.Println("alternately2:", string(b[i]))
		}
	}
}
