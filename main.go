package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home page"))
	})

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	router.HandleFunc("POST /hello/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("user_id")
		w.Write([]byte(fmt.Sprintf("hello user: %v", userId)))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("error while starting the server:", err)
	}

}
