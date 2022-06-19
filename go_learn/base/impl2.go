package bs

import "fmt"

func InterfaceImpl() {
	name := "San"
	s, err := New(name)
	if err != nil {
		fmt.Println("err:", err)
	}
	str := "点点"
	fmt.Println(s.Create(str))
	fmt.Println(s.Update(str))
	fmt.Println(s.Read(str))
	fmt.Println(s.Delete(str))
}
