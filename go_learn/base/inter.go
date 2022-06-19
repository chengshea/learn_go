package bs

import "fmt"

//电脑usb借口
type Usb interface {
	con()
	dis()
}

type Computer struct {
	dvd string
}

func (c Computer) Work(u Usb) {
	u.con()
	fmt.Println(c.dvd, "处理业务")
	u.dis()
}

type Phone struct {
	name string
}

func (p Phone) con() {
	fmt.Println("phone con ", p)
}

func (p Phone) dis() {
	fmt.Println("phone dis ", p)
}

func UseCom() {
	p := Phone{
		name: "ios",
	}
	c := Computer{
		dvd: "mac",
	}
	c.Work(p)
}
