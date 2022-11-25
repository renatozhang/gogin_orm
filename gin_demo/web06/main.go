package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//解析模板
	t := template.New("f.tmpl") //创建一个名字是f的模板对象，名字一定要与模板的名字能对应上
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("paese template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse tempalte failed, err:%v\n", err)
		return
	}
	name := "小王子"
	// 渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpl", tmplDemo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v", err)
		return
	}
}
