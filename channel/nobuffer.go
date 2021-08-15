package main

func main() {
	court := make(chan int)
	go test(court)
	court <- 1
}

func test(court chan int) {
	a := <-court
	println(a)
}
