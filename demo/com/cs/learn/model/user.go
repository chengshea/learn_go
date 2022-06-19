package model

type User struct {
	ID   int
	Name string
	Age  int
	Sex  bool //0 false;1 true
	//Like string
}

/*

grom TableName映射数据库实际表名info
如果不设置表名默认为结构体名user
*/
func (User) TableName() string {
	return "info"
}

func sex() string {

	return "女"
}
