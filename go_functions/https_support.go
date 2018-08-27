/*

Copyright () 2018

All rights reserved

FILE: https_support.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/8/27 上午11:22

*/

package main

import (
	"net/url"
	"net/http"
	"crypto/tls"
	"net/http/httputil"
	"fmt"
)

func main() {
	u, err := url.Parse("https://localhost:1087")
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(u),
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", dump)
}
