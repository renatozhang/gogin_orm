package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile("./template/hello.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, _ = fmt.Fprintln(w, string(html))
}

func main() {
	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v \n", err)
		return
	}
}
