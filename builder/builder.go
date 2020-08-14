package builder

import "fmt"

/**
建造者模式：指挥者角色Director, 抽象建造者Abstract Builder, 具体建造者Builder. 产品抽象组件。
	指挥者可以使用实现了抽象建造者的不同具体建造者对象，创建实现了产品抽象组件的不同产品
*/

// Director 指挥者
type Director struct {
	Builder Builder
}

// NewDirector 创建指挥者
func NewDirector(b Builder) *Director {
	return &Director{b}
}

// Builder 抽象建造者
type Builder interface {
	BuildPartA()
	BuildPartB()
	Construct()
}

// ActBuilder 具体建造者
type ActBuilder struct {
	Product Product
}

// BuildPartA 构建partA
func (a *ActBuilder) BuildPartA() {
	a.Product.makePartA()
}

// BuildPartB 构建partB
func (a *ActBuilder) BuildPartB() {
	a.Product.makePartB()
}

// AddFeture 添加特性
func (a *ActBuilder) AddFeture() {
	fmt.Println("ActBuiler 可以给产品添加额外特性")
}

// Construct 开始构建
func (a *ActBuilder) Construct() {
	fmt.Println("ActBuilder 开始构建。。。。。。")
	a.BuildPartA()
	a.BuildPartB()
	fmt.Println("ActBuilder 构建结束")
}

// BBuilder 具体建造者B
type BBuilder struct {
	Product Product
}

// BuildPartA 构建partA
func (b *BBuilder) BuildPartA() {
	b.Product.makePartA()
}

// BuildPartB 构建partB
func (b *BBuilder) BuildPartB() {
	b.Product.makePartB()
}

// GetProduct 获取所造产品
func (b *BBuilder) GetProduct() Product {
	return b.Product
}

// Construct 开始构建
func (b *BBuilder) Construct() {
	fmt.Println("BBuilder 开始构建。。。。。。")
	b.BuildPartA()
	b.BuildPartB()
	fmt.Println("BBuilder 构建结束")
}

// Product 产品抽象组件
type Product interface {
	makePartA()
	makePartB()
}

// AProduct A产品
type AProduct struct {
	ProductName string
}

// MakePartA make parta
func (a *AProduct) makePartA() {
	fmt.Println("正在创建产品：", a.ProductName, " 的A组件")
}

// MakePartB make partb
func (a *AProduct) makePartB() {
	fmt.Println("正在创建产品：", a.ProductName, " 的B组件")
}

// BProduct B产品
type BProduct struct {
	ProductName string
	LimitSale int
}

// MakePartA make parta
func (a *BProduct) makePartA() {
	fmt.Println("正在创建产品：", a.ProductName, " 的A组件")
}

// MakePartB make partb
func (a *BProduct) makePartB() {
	fmt.Println("正在创建产品：", a.ProductName, " 的B组件")
}

func (a *BProduct) limitSale(num int) {
	a.LimitSale = num
}

func (a *BProduct) getLimitSale() int {
	return a.LimitSale
}