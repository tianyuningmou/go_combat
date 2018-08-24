/*

Copyright () 2018

All rights reserved

FILE: handle_timeout.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/8/24 下午2:37

*/

package main

import (
	"testing"
	"net/http"
	"time"
	"encoding/json"
	"io"
	"net/http/httptest"
)

func TestClientTimeout(t *testing.T) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := map[string]interface{}{
			"id":    "12",
			"scope": "test-scope",
		}

		time.Sleep(100 * time.Millisecond)
		b, err := json.Marshal(d)
		if err != nil {
			t.Error(err)
		}
		io.WriteString(w, string(b))
		w.WriteHeader(http.StatusOK)
	})

	backend := httptest.NewServer(http.TimeoutHandler(handlerFunc, 20*time.Millisecond, "server timeout"))

	url := backend.URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("Request Error!", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err !=nil {
		t.Error("Response Error", err)
		return
	}

	defer resp.Body.Close()
}
