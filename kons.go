package main

import "fmt"

type listType func(selector selectorType) interface{}

// lisp car
func kar(x listType) interface{} {
	return x(func(a func(listType), b interface{}, c listType) interface{} {
		return b
	})
}

// lisp cdr
func kdr(x listType) interface{} {
	return x(func(a func(listType), b interface{}, c listType) interface{} {
		return c
	})
}

type selectorType func(func(listType), interface{}, listType) interface{}

func bons(kar interface{}) func(selectorType) interface{} {
	var kdr listType
	return func(selector selectorType) interface{} {
		return selector(func(x listType) {
			kdr = x
		}, kar, kdr)
	}
}

func setKdr(c func(selectorType) interface{}, x listType) {
	changer := c(func(f func(listType), i interface{}, l listType) interface{} {
		return f
	}).(func(listType)) // 通过 selector 获取 bons 中用来更改 kdr 的函数
	changer(x)          // 更改闭包中的值
}

// lisp cons
func kons(kar interface{}, kdr listType) listType {
	var a = bons(kar)
	setKdr(a, kdr)
	return a
}

func main() {
	a := kons(1, nil)
	b := kons(2, a)
	fmt.Println(kar(a))
	fmt.Println(kdr(a))
	fmt.Println(kar(b))
	fmt.Println(kdr(b))
}
