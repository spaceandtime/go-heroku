package main

import (
	"bufio"
	"net/http"
	"fmt"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	creq, err := http.ReadRequest(bufio.NewReader(r.Body))
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Transport.RoundTrip(creq)
	if err != nil {
		fmt.Printf("client error:%v\n", err)
		return
	}
	defer resp.Body.Close()
	if err := resp.Write(w); err != nil {
		fmt.Printf("Write error:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/http", proxyHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
