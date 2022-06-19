package bs

import (
	"encoding/json"
	"fmt"

	ms "github.com/mitchellh/mapstructure"
)

var (
	msg = []byte(`{
	"type": "UPDATE",
	"database": "blog",
	"table": "blogs",
	"data": [
			{
					"blogId": "100001",
					"title": "title",
					"content": "this is a blog",
					"uid": "1000012",
					"state": "1"
			}
	]}`)
)

func Js() {

	var event map[string]interface{}
	if err := json.Unmarshal(msg, &event); err != nil {
		panic(err)
	}

	fmt.Println(event)

	e := Event{}
	if err := json.Unmarshal(msg, &e); err != nil {
		panic(err)
	}
	var blogs []Blog2
	//弱解码 类型检查 数字字符串
	if err := ms.WeakDecode(e.Data, &blogs); err != nil {
		panic(err)
	}

	fmt.Println(blogs)

	ss(e)
}

func ss(e Event) {
	var blogs []Blog1
	if err := ms.Decode(e.Data, &blogs); err != nil {
		//panic(err)
		fmt.Println("Decode err:", err)
	}
	for k, v := range blogs {
		fmt.Printf("ss blogs[%d] %s\n", k, v)
		if k == 0 {
			fmt.Printf("blogs %s\n", v.Title)
		}
	}
	fmt.Println("ss", blogs)

}

func AutoGen() {
	e := AutoGenBlog{}
	if err := json.Unmarshal(msg, &e); err != nil {
		fmt.Println("Unmarshal err:", err)
	}
	if e.Data != nil {
		for k, v := range e.Data {
			fmt.Printf("AutoGenBlog %d ,%s\n", k, v.Title)
		}
		//空白符_
		for _, v := range e.Data {
			fmt.Printf("AutoGenBlog %s\n", v.Title)
		}
	}

}
