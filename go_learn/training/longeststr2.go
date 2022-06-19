package training

import (
	"fmt"
	"time"
)

func Plong() {
	stt := time.Now()
	str := "acabca"
	//simple(str)
	lengthOfLongestSubstring(str)
	defer func() {
		cost := time.Since(stt)
		fmt.Println(cost)
	}()
}

func simple(str string) {
	b := []byte(str)
	var dict [256]int //ASCII码256包含汉字
	st := 0
	ed := 0
	dict[0] = 0
	for i := 0; i < len(b); i++ {
		//
		if dict[b[i]] >= st {
			st = dict[b[i]]
		}
		if ed < i+1-st {
			ed = i + 1 - st
		}
		//
		dict[b[i]] = i + 1

	}
	fmt.Println(ed)
}

func lengthOfLongestSubstring(s string) int {
	n, l, p := 0, 0, -1
	substr := make([]rune, 0, len(s))
	for _, c := range s {
		p = rpos(substr, c)
		if p > -1 {
			l = len(substr)
			if n < l {
				n = l
			}
			//从怕p+1位开始截
			substr = append(substr[p+1:], c)
			//	fmt.Println("s", substr[p+1:], c)
		} else {
			substr = append(substr, c)
		}
	}

	if len(substr) > n {
		return len(substr)
	}
	//bca a b
	//fmt.Println(string(substr), string(substr[2:]), string(substr[:1]))
	return n
}

func rpos(runes []rune, r rune) int {
	for i, c := range runes {
		if c == r {
			return i
		}
	}
	return -1
}
