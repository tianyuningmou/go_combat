/*

Copyright () 2018

All rights reserved

FILE: goroutine_leak.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/8/22 下午5:08

*/

package main

import (
	"math/rand"
	"time"
	"fmt"
	"runtime"
)

// 发送到一个没有接收者的channel

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func queryAll() int {
	ch := make(chan int)
	go func() {ch <- query() }()
	go func() {ch <- query() }()
	go func() {ch <- query() }()
	return <-ch
}

// 从没有发送者的channel中接收数据

func nil_channel()  {
	var ch chan struct{}
	ch <- struct{}{}
}

func main()  {
	for i := 0; i < 4; i++ {
		queryAll()
		fmt.Printf("#gotourines: %d \n", runtime.NumGoroutine())
	}

	nil_channel()
}
