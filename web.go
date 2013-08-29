package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	for {
		fmt.Fprintln(res, " ")
		time.Sleep(10 * time.Second)
	}
}
