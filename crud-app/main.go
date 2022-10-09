package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", PostController)
	http.HandleFunc("/posts/", PostController)
	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
