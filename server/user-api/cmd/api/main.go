package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HelloWorldaaa"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("サーバー起動に失敗しました:", err)
	}
}
