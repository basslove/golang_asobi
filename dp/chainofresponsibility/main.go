package main

import (
	"fmt"
	"strconv"
)

type Trouble struct {
	number int
}

func (t *Trouble) getNumber() int {
	return t.number
}

type Support interface {
	Resolve(trouble Trouble) bool
	Handle(support Support, trouble Trouble) string
}

type initSupport struct {
	Support
	name string
	next Support
}

func (is *initSupport) SetNext(next Support) {
	is.next = next
}

func (is *initSupport) Handle(support Support, trouble Trouble) string {
	if support.Resolve(trouble) {
		return is.done(trouble)
	} else if is.next != nil {
		return is.next.Handle(is.next, trouble)
	} else {
		return is.fail(trouble)
	}
}

func (is *initSupport) done(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " is resolved by " + is.name
}

func (is *initSupport) fail(trouble Trouble) string {
	return "trouble:" + strconv.Itoa(trouble.getNumber()) + " cannot be resolved"
}

type noSupport struct {
	*initSupport
}

func (ns *noSupport) Resolve(trouble Trouble) bool {
	return false
}

func NewNoSupport(name string) *noSupport {
	return &noSupport{&initSupport{name: name}}
}

type limitSupport struct {
	*initSupport
	limit int
}

func (ls *limitSupport) Resolve(trouble Trouble) bool {
	if trouble.getNumber() < ls.limit {
		return true
	} else {
		return false
	}
}

func NewLimitSupport(name string, limit int) *limitSupport {
	return &limitSupport{&initSupport{name: name}, limit}
}

func main() {
	a := NewNoSupport("A")
	b := NewLimitSupport("B", 2)
	c := NewLimitSupport("C", 3)
	a.SetNext(b)
	b.SetNext(c)

	fmt.Println(a.Handle(a, Trouble{1}))
	fmt.Println(a.Handle(a, Trouble{2}))
	fmt.Println(a.Handle(a, Trouble{3}))
}
