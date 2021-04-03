package main

import "fmt"

type ListNode struct {
	Val  *int
	Next *ListNode
}

func makeList(data []int) *ListNode {
	var head, prev *ListNode
	for i := range data {
		node := &ListNode{&data[i], nil}
		if prev != nil {
			prev.Next = node
		} else {
			head = node
		}
		prev = node
	}
	return head
}

func printList(head *ListNode) {
	for {
		if head.Val != nil {
			fmt.Println(*head.Val)
		}
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
}

func forever(x *ListNode) int {
	return forever(x)
}

func trap(x interface{}) int {
	panic(1)
}

func main() {
	fmt.Println(length(makeList([]int{1, 2, 3, 4, 5, 6})))
	firstLevel()
	secondLevel()
	thirdLevel()
	fourthLevel()
	fifthLevel()
	sixthLevel()
}

func length(node *ListNode) int {
	if node == nil {
		return 0
	} else {
		return 1 + length(node.Next)
	}
}

// length(n) = 1 + length(n-1)
func firstLevel() {
	// length0, 只能计算含有 0 个元素的链表的长度
	result := func(node *ListNode) int {
		if node == nil {
			return 0
		} else {
			return 1 + forever(node.Next)
		}
	}(nil)
	fmt.Println(result)

	data1 := makeList([]int{1})
	// length1, 只能计算含有 1 个元素的链表的长度
	result = func(node *ListNode) int {
		if node == nil {
			return 0
		} else {
			// 1 + length0
			return 1 + func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + forever(node.Next)
				}
			}(node.Next)
		}
	}(data1)
	fmt.Println(result)

	// length2, 只能计算含有 2 个元素的链表的长度
	data2 := makeList([]int{1, 2})
	result = func(node *ListNode) int {
		if node == nil {
			return 0
		} else {
			// 1 + length1
			return 1 + func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					// 1 + length0
					return 1 + func(node *ListNode) int {
						if node == nil {
							return 0
						} else {
							return 1 + forever(node.Next)
						}
					}(node.Next)
				}
			}(node.Next)
		}
	}(data2)
	fmt.Println(result)

	// length n ...
}

type lengthType func(*ListNode) int

// make-length(n) = make-length(make-length(n-1))
func secondLevel() {
	// make length 0 = make-length(forever)
	result := func(f lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 0
			} else {
				return 1 + f(node.Next)
			}
		}
	}(forever)(nil)
	fmt.Println(result)

	data1 := makeList([]int{1})
	// make length 1 = make-length(make length 0)
	result = func(f lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 0
			} else {
				return 1 + f(node.Next)
			}
		}
	}(
		func(f lengthType) lengthType {
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(node.Next)
				}
			}
		}(forever),
	)(data1)
	fmt.Println(result)

	data2 := makeList([]int{1, 2})
	// make length 2 = make-length(make length 1)
	result = func(f lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 0
			} else {
				return 1 + f(node.Next)
			}
		}
	}(
		func(f lengthType) lengthType { // make length 1
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(node.Next)
				}
			}
		}(
			func(f lengthType) lengthType { // make length 0
				return func(node *ListNode) int {
					if node == nil {
						return 0
					} else {
						return 1 + f(node.Next)
					}
				}
			}(forever)),
	)(data2)
	fmt.Println(result)

	// make length 3 = make-length(make-length(make-length(make-length forever))))

	// make length n ...
}

type lengthTypeMaker func(f lengthType) lengthType

// make length n = make-length(make-length(...(make-length forever))...)
func thirdLevel() {
	// length 0
	result := func(f lengthTypeMaker) lengthType {
		return f(forever)
	}(
		func(f lengthType) lengthType {
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(node.Next)
				}
			}
		},
	)(nil)
	fmt.Println(result)

	data1 := makeList([]int{1})
	// length 1
	result = func(f lengthTypeMaker) lengthType {
		return f(f(forever))
	}(
		func(f lengthType) lengthType {
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(node.Next)
				}
			}
		},
	)(data1)
	fmt.Println(result)

	data2 := makeList([]int{1, 2})
	// length 2
	result = func(f lengthTypeMaker) lengthType {
		return f(f(f(forever)))
	}(
		func(f lengthType) lengthType {
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(node.Next)
				}
			}
		},
	)(data2)
	fmt.Println(result)
}

type lengthTypeMaker2 func(maker lengthTypeMaker2) lengthType

// recursive
func fourthLevel() {
	data := makeList([]int{1, 2, 3, 4})
	result := func(f lengthTypeMaker2) lengthType {
		return f(f)
	}(
		func(f lengthTypeMaker2) lengthType {
			return func(node *ListNode) int {
				if node == nil {
					return 0
				} else {
					return 1 + f(f)(node.Next) // 递归开始
				}
			}
		},
	)(data)
	fmt.Println(result)
}

// 再抽象一层，将 f(f) 提取到外层
func fifthLevel() {
	data := makeList([]int{1, 2, 3, 4, 5})
	result := func(f lengthTypeMaker2) lengthType {
		return f(f)
	}(
		func(f lengthTypeMaker2) lengthType {
			return func(g lengthType) lengthType {
				return func(node *ListNode) int {
					if node == nil {
						return 0
					} else {
						return 1 + g(node.Next) // 递归开始
					}
				}
			}(func(node *ListNode) int {
				return f(f)(node)
			})
		},
	)(data)
	fmt.Println(result)
}

// 再抽象一层，将计算逻辑提取到外层
func sixthLevel() {
	data := makeList([]int{1, 2, 3, 4, 5, 6})
	result := func(le lengthTypeMaker) lengthType {
		return func(f lengthTypeMaker2) lengthType {
			return f(f)
		}(
			func(f lengthTypeMaker2) lengthType {
				return le(func(node *ListNode) int {
					return f(f)(node)
				})
			},
		)
	}(func(g lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 0
			} else {
				return 1 + g(node.Next) // 递归开始
			}
		}
	})(data)
	fmt.Println(result)

	// 提取递归结构
	y := func(le lengthTypeMaker) lengthType {
		return func(f lengthTypeMaker2) lengthType {
			return f(f)
		}(
			func(f lengthTypeMaker2) lengthType {
				return le(func(node *ListNode) int {
					return f(f)(node)
				})
			},
		)
	}

	// 求和
	result = y(func(f lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 0
			} else {
				return *node.Val + f(node.Next) // 递归开始
			}
		}
	})(data)
	fmt.Println(result)

	// 求乘积
	result = y(func(f lengthType) lengthType {
		return func(node *ListNode) int {
			if node == nil {
				return 1
			} else {
				return *node.Val * f(node.Next) // 递归开始
			}
		}
	})(data)
	fmt.Println(result)
}
