package main

import "fmt"
//go语言的函数式编程
func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
func main() {
	s1 := test(func() int { return 100 })
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20)
	println(s1, s2)
}
