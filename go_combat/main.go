/*

Copyright () 2018

All rights reserved

FILE: go_combat.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/7/25 下午1:06

*/

package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"html/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	// 解析url传递的参数，对于POST请求则解析响应包的主体
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello everyone!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("go_combat/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main()  {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
