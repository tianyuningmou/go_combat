package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.baidu.com")
	if err == nil {
		defer req.Body.Close()
		content, err := ioutil.ReadAll(req.Body)
		if err == nil {
			fmt.Printf(string(content))
		}
	}
}