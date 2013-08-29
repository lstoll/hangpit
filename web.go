package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", hang)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hang(res http.ResponseWriter, req *http.Request) {
	for {
		fmt.Fprintln(res, "\n")
		time.Sleep(5 * time.Second)
		res.(http.Flusher).Flush()
	}
}
