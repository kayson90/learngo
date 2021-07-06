package main

import "fmt"

func main() {
	//var i int = 10
	//fmt.Println("i的地址=", &i)
	//
	//var ptr *int = &i;
	//fmt.Printf("ptr=%v\n", ptr)
	//fmt.Printf("ptr 的地址=%v", &ptr)
	//fmt.Printf("ptr 指向的值=%v", *ptr)

	var num int = 9
	fmt.Printf("num的地址是=%v\n",&num)
	var ptr *int
	ptr = &num
	*ptr=10
	fmt.Println("num=",num)
}
