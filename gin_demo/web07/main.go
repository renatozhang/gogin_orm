package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./index.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, err:%v", err)
		return
	}

	msg := "这是index页面"

	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./base.tmpl", "./home.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, err:%v", err)
		return
	}

	msg := "小王子"

	t.Execute(w, msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
