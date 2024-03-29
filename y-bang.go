package main

import "fmt"

type lengthTypeYbang func(list []interface{}) int
type L_Type func(length lengthTypeYbang) lengthTypeYbang

// 递归版本
func length1(list []interface{}) int {
	if len(list) == 0 {
		return 0
	} else {
		return 1 + length1(list[1:])
	}
}

// 赋值版本
func length2(list []interface{}) int {
	h := func(list []interface{}) int { return 0 }
	h = func(list []interface{}) int {
		if len(list) == 0 {
			return 0
		} else {
			return 1 + h(list[1:])
		}
	}
	return h(list)
}

// 赋值版本，x 的初始值是什么无所谓，只要满足引用自身即可
func length3(list []interface{}) int {
	// lisp 里面没有变量声明，所以需要定义 (let ((h lambda (l) (quote ())))...)
	var h lengthTypeYbang
	// 引用自身
	h = func(list []interface{}) int {
		if len(list) == 0 {
			return 0
		} else {
			return 1 + h(list[1:])
		}
	}
	return h(list)
}

// 消去赋值部分
func length4(list []interface{}) int {

	// 接收一个 length 类型的函数，返回包装了 length 的函数
	L := func(length lengthTypeYbang) lengthTypeYbang {
		return func(list []interface{}) int {
			if len(list) == 0 {
				return 0
			} else {
				return 1 + length(list[1:])
			}
		}
	}

	// lisp 里面没有变量声明，所以需要定义 (let ((h lambda (l) (quote ())))...)
	var h lengthTypeYbang
	// 新的 length 函数
	// h 的值只变化一次，在执行完 L 后，h 的值才被确定，作为 L 参数的函数中的 h 的值也才被确定
	h = L(func(arg []interface{}) int {
		return h(arg)
	})

	/* h 相当于
	h = func(list []interface{}) int {
		if len(list) == 0 {
			return 0
		} else {
			return 1 + func(arg []interface{}) int {
				return h(arg)
			}(list[1:])
		}
	}
	*/
	return h(list)

}

// 使用 L 定义 Y-bang
func length5(list []interface{}) int {

	// 接收一个 length 类型的函数，返回包装了 length 的函数
	L := func(length lengthTypeYbang) lengthTypeYbang {
		return func(list []interface{}) int {
			if len(list) == 0 {
				return 0
			} else {
				return 1 + length(list[1:])
			}
		}
	}

	yBang := func(L L_Type) lengthTypeYbang {
		var h lengthTypeYbang
		h = L(func(list []interface{}) int {
			return h(list)
		})
		return h
	}

	length := yBang(L)

	return length(list)
}

func main() {
	fmt.Println(length1([]interface{}{1, 2, 3, 4}))
	fmt.Println(length2([]interface{}{1, 2, 3, 4}))
	fmt.Println(length3([]interface{}{1, 2, 3, 4}))
	fmt.Println(length4([]interface{}{1, 2, 3, 4}))
	fmt.Println(length5([]interface{}{1, 2, 3, 4}))
}
