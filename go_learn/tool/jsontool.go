package tool

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int                    `json:"code"`
	Message string                 ` json:"msg"`
	Data    map[string]interface{} ` json:"data"`
}

func Tojson(res *Result) {
	jsons, errs := json.Marshal(res)
	exption("to json.Marshal error:", errs)
	fmt.Println("json:", string(jsons))
}

func Dejson(jsons []byte) map[string]interface{} {
	ser := make(map[string]interface{})
	errs := json.Unmarshal(jsons, &ser)
	exption("json.Unmarshal error:", errs)
	fmt.Println("fromtjson:", ser)
	return ser
}

func exption(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
	}
}
