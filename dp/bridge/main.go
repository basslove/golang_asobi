package main

import (
	"fmt"
)

type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

var _ DisplayImpl = &StringDisplayImpl{}

type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) *Display {
	return &Display{impl: impl}
}

func (d *Display) open() {
	d.impl.RawOpen()
}

func (d *Display) print() {
	d.impl.RawPrint()
}

func (d *Display) close() {
	d.impl.RawClose()
}

func (d *Display) display() {
	d.open()
	d.print()
	d.close()
}

type CountDisplay struct {
	d Display
}

func NewCountDisplay(display Display) *CountDisplay {
	return &CountDisplay{d: display}
}

func (cd *CountDisplay) multiDisplay(times int) {
	cd.d.open()
	for range make([]int, times) {
		cd.d.print()
	}
	cd.d.close()
}

type StringDisplayImpl struct {
	s     string
	width int
}

func NewStringDisplayImpl(s string) *StringDisplayImpl {
	return &StringDisplayImpl{s: s, width: len(s)}
}

func (sd *StringDisplayImpl) RawOpen() {
	sd.printLine()
}

func (sd *StringDisplayImpl) RawPrint() {
	fmt.Printf("| %v |\n", sd.s)
}

func (sd *StringDisplayImpl) RawClose() {
	sd.printLine()
}

func (sd *StringDisplayImpl) printLine() {
	fmt.Print("+")
	for range make([]int, sd.width) {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func main() {
	fmt.Println("bridge")

	d1 := NewDisplay(NewStringDisplayImpl("hello japan"))
	d2 := NewCountDisplay(*NewDisplay(NewStringDisplayImpl("hello world")))
	d3 := NewCountDisplay(*NewDisplay(NewStringDisplayImpl("hello universe")))
	d1.display()
	d2.d.display()
	d3.d.display()
	d3.multiDisplay(5)
}
