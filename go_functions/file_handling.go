/*

Copyright () 2018

All rights reserved

FILE: file_handling.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/9/5 下午3:44

*/

package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs(".")
	data, err := ioutil.ReadFile(path + "/go_functions/test.txt")
	if err != nil {
		fmt.Println("File reading err", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
