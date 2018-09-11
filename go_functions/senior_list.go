/*

Copyright () 2018

All rights reserved

FILE: senior_list.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/9/11 下午5:08

*/

package main

import "fmt"

func main() {
	var times [5][0]int
	for range times{
		fmt.Println("Hello")
	}

	var a = [...]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("a[%d]: %d\n", i, v)
	}

	var b = &a
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
}
