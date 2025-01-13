package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello from Hell!!\n")
	}

	http.HandleFunc("/", helloHandler)

	log.Println("listening at port 8080")

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) // exit 1のときにログ出力される
}
