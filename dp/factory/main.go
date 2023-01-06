package main

import (
	"fmt"
)

type Product interface {
	Use()
}

type FactoryMethod interface {
	CreateProduct(owner string) *IDCard
	RegisterProduct(owner string)
}

type Factory struct {
	method FactoryMethod
}

func NewFactory(method FactoryMethod) *Factory {
	return &Factory{method: method}
}

func (f *Factory) Create(owner string) Product {
	p := f.method.CreateProduct(owner)
	f.method.RegisterProduct(owner)
	return p
}

type IDCard struct {
	owner string
}

func NewIDCard(owner string) *IDCard {
	return &IDCard{owner: owner}
}

func (i *IDCard) Use() {
	fmt.Printf("%vを使います\n", i.getOwner())
}

func (i *IDCard) toString() {
	fmt.Printf("[IDCard:%v]", i.owner)
}

func (i *IDCard) getOwner() string {
	return i.owner
}

type IDCardFactory struct{}

func NewIDCardFactory() *IDCardFactory {
	return &IDCardFactory{}
}

func (f *IDCardFactory) CreateProduct(owner string) *IDCard {
	return NewIDCard(owner)
}

func (f *IDCardFactory) RegisterProduct(owner string) {
	fmt.Printf("%vを登録しました\n", owner)
}

func main() {
	fmt.Println("factory")

	var fm FactoryMethod = NewIDCardFactory()
	factory := NewFactory(fm)
	card1 := factory.Create("aiueo")
	card2 := factory.Create("12345")
	card3 := factory.Create("#$%^&")

	card1.Use()
	card2.Use()
	card3.Use()
}
