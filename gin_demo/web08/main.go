package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, err:%v", err)
		return
	}

	msg := "这是index页面"

	t.Execute(w, msg)
}

func xss(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, err:%v", err)
		return
	}
	jsStr := `<script>alert('嘿嘿嘿')</script>`
	data := map[string]interface{}{"str1": jsStr, "str2": jsStr}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
