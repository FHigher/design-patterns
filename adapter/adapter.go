package adapter

import "fmt"

/**
适配器模式：通过适配器让原本不兼容的接口，重新兼容。分为类适配器，对象适配器
*/

// Target 目标接口
type Target interface {
	request()
}

// Adaptee 已有的不兼容的接口
type Adaptee struct{}

func (a *Adaptee) doRequest() {
	fmt.Println("This is a request")
}

// ClassAdapter 类适配器
type ClassAdapter struct{}

func (c *ClassAdapter) request() {
	adaptee := &Adaptee{}
	adaptee.doRequest()
}

// ObjAdapter 对象适配器
type ObjAdapter struct {
	adaptee *Adaptee
}

// NewObjAdapter 创建对象适配器
func NewObjAdapter(adaptee *Adaptee) *ObjAdapter {
	return &ObjAdapter{adaptee}
}

func (o *ObjAdapter) request() {
	o.adaptee.doRequest()
}
