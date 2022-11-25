package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template failed,err:", err)
		return
	}
	// 渲染模板
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    0,
	}

	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    0,
	}

	hobbyList := []string{"篮球", "足球", "双色球"}
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("render template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
