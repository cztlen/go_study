package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
