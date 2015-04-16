package main

import (
	"bufio"
	"net/http"
	"fmt"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	creq, err := http.ReadRequest(bufio.NewReader(r.Body))
	if err != nil {
		s.HTTPError("ReadRequest: %s", err.Error())
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

	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
