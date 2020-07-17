package main

import (
	"fmt"
	"net/http"
)

type ServeString string

func (s ServeString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/hello", hello)

	http.Handle("/bye", ServeString("Bye"))
	// リクエスト方式をPOSTに絞る
	http.HandleFunc("/only_post", handleOnlyPost)
	// htmlファイルを配信
	http.Handle("/", http.FileServer(http.Dir("./contents")))

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from Go.")
}

func handleOnlyPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("Only Post available"))
		return
	}

	fmt.Fprint(w, "OK")
}
