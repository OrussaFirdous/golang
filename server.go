package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hanlder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
	fmt.Println("URL.path ", r.URL.Path)
}
