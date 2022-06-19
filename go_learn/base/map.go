package bs

import (
	"encoding/json"
	"fmt"
	"go_learn/tool"
)

func Map2() {
	var res tool.Result
	res.Code = 200
	res.Message = "success"
	res.Data = map[string]interface{}{
		"user": "撒三",
		"age":  8,
		"like": []string{"游戏", "电影"},
	}
	tool.Tojson((*tool.Result)(&res))
}

func Map1() {

	student := map[string]string{
		"xk":  "爪哇",
		"sex": "女",
	}
	fmt.Println("s:", student)
	delete(student, "xk")
	student["age"] = "21"
	fmt.Println("del xk:", student)

	res := make(map[string]interface{})
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"user": "利索",
		"age":  88,
		"like": []string{"游戏", "电影"},
	}

	fmt.Println("res:", res)
	jsons, errs := json.Marshal(res)
	errf(errs)
	fmt.Println("json res:", string(jsons))
	// ser := make(map[string]interface{})
	// errs = json.Unmarshal([]byte(jsons), &ser)
	// errf(errs)
	// fmt.Println("data ser:", ser)
	res = tool.Dejson(jsons)
	//Decode(res,&)

	str := tool.MD5("cs")
	fmt.Println("str:", str)
	str = tool.MD5secert("cs", "")
	fmt.Println("1 MD5secert:", str)
	str = tool.MD5secert("cs", "ssl")
	fmt.Println("2 MD5secert:", str)
}

func errf(errs error) {
	if errs != nil {
		fmt.Println("error:", errs)
	}
}
