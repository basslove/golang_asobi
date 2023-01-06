package main

import "fmt"

type AbstractMethod interface {
	Open()
	Print()
	Close()
}

type AbstractDisplay struct {
	method AbstractMethod
}

func (am *AbstractDisplay) Display() {
	am.method.Open()
	for range make([]int, 5) {
		am.method.Print()
	}
	am.method.Close()
}

type CharDisplay struct {
	ch string
}

func NewCharDisplay(ch string) *CharDisplay {
	return &CharDisplay{ch: ch}
}

func (cd *CharDisplay) Open() {
	fmt.Printf("<<")
}

func (cd *CharDisplay) Print() {
	fmt.Printf(cd.ch)
}

func (cd *CharDisplay) Close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	ch    string
	width int
}

func NewStringDisplay(ch string) *StringDisplay {
	return &StringDisplay{ch: ch, width: len(ch)}
}

func (cd *StringDisplay) Open() {
	cd.printLine()
}

func (cd *StringDisplay) Print() {
	fmt.Printf("| %v |\n", cd.ch)
}

func (cd *StringDisplay) Close() {
	cd.printLine()
}

func (cd *StringDisplay) printLine() {
	fmt.Printf("*")

	for range make([]int, cd.width) {
		fmt.Printf("-")
	}

	fmt.Println("*")
}

func main() {
	fmt.Println("template method")
	d1 := &AbstractDisplay{method: NewCharDisplay("H")}
	d1.Display()
	d2 := AbstractDisplay{method: NewStringDisplay("hello world")}
	d2.Display()
}
