package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var a = 1

func main() {
	fmt.Println("this is a golang service")
	http.HandleFunc("/panic", PanicHandler)
	http.HandleFunc("/test", TestHandler)
	http.ListenAndServe(":8765", nil)
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	// 测试守护进程
	a := make(map[int]int)
	go func() {
		for i := 0; i < 100; i++ {
			a[i] = i
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			a[i] = i
		}
	}()
	fmt.Println(111)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	a++
	fmt.Fprintf(w, strconv.Itoa(a))
}
