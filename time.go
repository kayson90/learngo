package main

import (
	"fmt"
	"time"
)

type user struct {
	id      int64
	right   bool
	text    string
	endTime int64
}

func main() {
	stu := &user{
		id:      1,
		right:   true,
		text:    "no reason",
		endTime: 1625135600000,
	}

	if stu.right && stu.endTime > 0 && stu.endTime < time.Now().UnixNano()/1e6 {
		stu.right = false
		stu.text = "live auth expired"
	}
	msg := fmt.Sprintf("%+v:%+v", stu.id, *stu)
	println(msg)
}
