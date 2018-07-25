/*

Copyright () 2018

All rights reserved

FILE: custom_web.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/7/25 下午3:40

*/

package main

import (
	"net/http"
	"fmt"
)

type MyMux struct {

}

func (p *MyMux) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloBody(w, r)
		return
	}
	http.NotFound(w, r)
	return 
}

func sayHelloBody(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello everyone")
}

func main()  {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
