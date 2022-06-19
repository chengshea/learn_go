package bs

import (
	"fmt"
	"sync"
)

func m() {
	val := make(map[int]int)
	for i := 0; i < 1000; i++ {
		val[i] = i
	}
	fmt.Println("m len:", len(val))
}

/*
并发安全
*/
func sm() {
	val := sync.Map{}
	for i := 0; i < 1000; i++ {
		val.Store(i, fmt.Sprintf("%d%s", i, "s"))
	}
	len := 0
	val.Range(func(key, value any) bool {
		len++
		if len == 500 {
			fmt.Println(val.Load(len))
			fmt.Println(val.LoadOrStore(len, "hhhahhaha"))
		}
		return true
	})

	fmt.Println("sm len:", len)
}

func wg() {
	var wg sync.WaitGroup
	wg.Add(2)

	go handler1(&wg)
	go handler2(&wg)
	wg.Wait()
}

func handler2(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done() //相当于是Add(-1)
	fmt.Println("2 ", "业务")
}

//当第一个子任务执行失败时，通知其他子任务停止运行，wg是无法满足的
func handler1(waitGroup *sync.WaitGroup) {
	panic("unimplemented")
	//defer waitGroup.Done()
}

func SyncMap() {
	m()
	sm()
	wg()
}
