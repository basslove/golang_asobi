package main

import (
	"fmt"
)

type Product interface {
	Use(s string)
	CreateCopy() Product
}

type Manager struct {
	showcase map[string]Product
}

func NewManager() *Manager {
	return &Manager{showcase: make(map[string]Product)}
}

func (m *Manager) Register(name string, product Product) {
	m.showcase[name] = product
}

func (m *Manager) Create(protoTypeName string) Product {
	p := m.showcase[protoTypeName]
	return p.CreateCopy()
}

type MessageBox struct {
	decoChar string
}

func NewMessageBox(decoChar string) *MessageBox {
	return &MessageBox{decoChar: decoChar}
}

func (mb *MessageBox) Use(s string) {
	decoLen := 1 + len(s) + 1
	for range make([]int, decoLen) {
		fmt.Printf(mb.decoChar)
	}

	fmt.Println()
	fmt.Printf("%v %v %v\n", mb.decoChar, s, mb.decoChar)
	for range make([]int, decoLen) {
		fmt.Printf(mb.decoChar)
	}
	fmt.Println()
}

func (mb *MessageBox) CreateCopy() Product {
	return NewMessageBox(mb.decoChar)
}

type UnderlinePen struct {
	ulChar string
}

func NewUnderlinePen(ulChar string) *UnderlinePen {
	return &UnderlinePen{ulChar: ulChar}
}

func (mb *UnderlinePen) Use(s string) {
	uLen := len(s)

	fmt.Println(s)
	for range make([]int, uLen) {
		fmt.Printf(mb.ulChar)
	}
	fmt.Println()
}

func (mb *UnderlinePen) CreateCopy() Product {
	return NewUnderlinePen(mb.ulChar)
}

func main() {
	fmt.Println("prototype")

	manager := NewManager()
	uPen := NewUnderlinePen("-")
	mBox := NewMessageBox("*")
	sBox := NewMessageBox("/")

	manager.Register("strong message", uPen)
	manager.Register("warning box", mBox)
	manager.Register("slash box", sBox)

	var p1 Product = manager.Create("strong message")
	p1.Use("hello world")

	var p2 Product = manager.Create("warning box")
	p2.Use("hello world")

	var p3 Product = manager.Create("slash box")
	p3.Use("hello world")
}
