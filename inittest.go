package main

import "fmt"

var age = test()

func test() int{
	fmt.Println("test()")
	return 90
}
func init(){
	fmt.Println("初始化…")
}

func main() {
	fmt.Println("main()…age=",age)
}
