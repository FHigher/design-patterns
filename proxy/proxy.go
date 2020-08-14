package proxy

import "fmt"

/*
	代理模式：在访问者与真实对象间做转发，可以增添特性。
*/

// Obj obj
type Obj interface {
	dispaly(string)
}

// ActualObj actual obj
type ActualObj struct {
	Name string
}

func (a *ActualObj) dispaly(name string) {
	a.Name = name
	fmt.Println("真实对象 ", a.Name)
}

// Pobj 代理对象
type Pobj struct {
	*ActualObj
	count int // 统计真实对象被访问的次数
}

func (p *Pobj) dispaly(name string) {
	p.ActualObj = &ActualObj{}
	p.ActualObj.dispaly(name)
	p.count++
}

func (p *Pobj) getCount() {
	fmt.Println("真实对象被访问了 ", p.count)
}
