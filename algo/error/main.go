package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func newMyError(s string) *MyError {
	err := new(MyError)
	err.msg = s
	return err
}

func (err *MyError) Error() string {
	return err.msg
}

// error.new
func no1(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("fact : domain error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}

// my-error.new
func no2(n int) (int, error) {
	if n < 0 {
		return 0, newMyError("fact : domain my error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}

// panic
func no3(n int) (int, error) {
	if n < 0 {
		panic("fact : domain panic error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}

// defer
// defer 文の指定は、ランタイムエラーで関数の実行が中断された場合にも有効
func no4() {
	f0 := func() {
		panic("oops!")
	}

	f1 := func() {
		defer fmt.Println("bar end")
		fmt.Println("bar start!")
		f0()
	}
	defer fmt.Println("foo end")
	fmt.Println("foo start!")
	defer f1()
}

func main() {
	fmt.Println("error")

	//// no1
	//for x := 10; x >= -1; x-- {
	//	v, err := no1(x)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(v)
	//	}
	//}

	//// no2
	//for x := 10; x >= -1; x-- {
	//	v, err := no2(x)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(v)
	//	}
	//}

	//// no3
	//for x := 10; x >= -1; x-- {
	//	v, err := no3(x)
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(v)
	//	}
	//}

	no4()
}
