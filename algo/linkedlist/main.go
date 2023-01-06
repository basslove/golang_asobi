package main

import "fmt"

type Cell struct {
	item int
	next *Cell
}

func NewCell(x int, cp *Cell) *Cell {
	nc := new(Cell)
	nc.item, nc.next = x, cp

	return nc
}

type List struct {
	top *Cell
}

func NewList() *List {
	l := new(List)
	l.top = new(Cell)

	return l
}

func (c *Cell) nthCell(n int) *Cell {
	i := -1

	for c != nil {
		if i == n {
			return c
		}
		i++
		c = c.next
	}

	return nil
}

func (l *List) nth(n int) (int, bool) {
	cp := l.top.nthCell(n)
	if cp == nil {
		return 0, false
	}

	return cp.item, true
}

func (l *List) insertNth(n, x int) bool {
	cp := l.top.nthCell(n - 1)
	if cp == nil {
		return false
	}

	cp.next = NewCell(x, cp.next)

	return true
}

func (l *List) deleteNth(n int) bool {
	cp := l.top.nthCell(n - 1)
	if cp == nil || cp.next == nil {
		return false
	}

	cp.next = cp.next.next

	return true
}

func (lst *List) isEmpty() bool {
	return lst.top.next == nil
}

func (lst *List) printList() {
	cp := lst.top.next
	for ; cp != nil; cp = cp.next {
		fmt.Print(cp.item, " ")
	}
	fmt.Println("")
}

func main() {
	a := NewList()

	for i := 0; i < 4; i++ {
		fmt.Println(a.insertNth(i, i))
	}
	a.printList()

	for i := 0; i < 5; i++ {
		n, ok := a.nth(i)
		fmt.Println(n, ok)
	}

	for !a.isEmpty() {
		a.deleteNth(0)
		a.printList()
	}
}
