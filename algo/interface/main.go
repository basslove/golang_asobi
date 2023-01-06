package main

import "fmt"

func sumOfInt(ary []interface{}) int {
	sum := 0
	for _, x := range ary {
		v, ok := x.(int)
		if ok {
			sum += v
		}
	}
	return sum
}

// 実数の合計値を求める
func sumOfFloat(ary []interface{}) float64 {
	sum := 0.0
	for _, x := range ary {
		v, ok := x.(float64)
		if ok {
			sum += v
		}
	}
	return sum
}

type Num interface {
	number()
}

type Int int

func (n Int) number() {}

type Real float64

func (n Real) number() {}

func sumOfNum(ary []Num) (Int, Real) {
	var sumi Int = 0
	var sumr Real = 0.0
	for _, x := range ary {
		switch v := x.(type) {
		case Int:
			sumi += v
		case Real:
			sumr += v
		}
	}
	return sumi, sumr
}

type Foo struct {
	a int
}

type FooI interface {
	getA() int
}

func (p *Foo) getA() int { return p.a }

type Bar struct {
	b int
}

type BarI interface {
	getB() int
}

func (p *Bar) getB() int { return p.b }

type Baz struct {
	Foo
	Bar
}

type BazI interface {
	FooI
	BarI
}

func main() {
	a := []interface{}{1, 1.1, "abc", 2, 2.2, "def", 3, 3.3}
	fmt.Println(sumOfInt(a))
	fmt.Println(sumOfFloat(a))

	var ary []Num = []Num{
		Int(1), Real(1.1), Int(2), Real(2.2), Int(3), Real(3.3),
	}
	r1, r2 := sumOfNum(ary)
	fmt.Println(r1, r2)

	x := []FooI{
		&Foo{1}, &Foo{2}, &Baz{},
	}
	y := []BarI{
		&Bar{10}, &Bar{20}, &Baz{},
	}
	z := []BazI{
		&Baz{}, &Baz{Foo{1}, Bar{2}}, &Baz{Foo{3}, Bar{4}},
	}
	for i := 0; i < 3; i++ {
		fmt.Println(x[i].getA())
		fmt.Println(y[i].getB())
		fmt.Println(z[i].getA())
		fmt.Println(z[i].getB())
	}
}
