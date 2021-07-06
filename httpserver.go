package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	go http.ListenAndServe(":8080", mux)
	fmt.Println(<-sig)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
