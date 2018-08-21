/*

Copyright () 2018

All rights reserved

FILE: copy_tail.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/8/21 下午4:37

*/

package go_functions

import (
	"os"
	"github.com/fsnotify/fsnotify"
	"bufio"
	"io"
	"fmt"
)

/*
  defer关键字
    1、defer nil
	    如果一个延迟函数被赋值为nil，运行时的panic异常会发生在外围函数执行结束后而不是defer的函数被调用的时候
	2、在循环中使用defer
		不要在循环中使defer，可能会导致延迟函数不断堆积到延迟调用栈中。
	3、延迟调用含有闭包的函数
		调用含有闭包的函数，可能会出现错误
	4、在执行块中使用defer
	5、延迟方法的坑
*/

func follow(filename string) error {
	file, _ := os.Open(filename)
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	_ = watcher.Add(filename)

	r := bufio.NewReader(file)
	for {
		by, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Print(string(by))
		if err != io.EOF {
			continue
		}
		if err = waitForChange(watcher); err != nil {
			return err
		}
	}
}

func waitForChange(w *fsnotify.Watcher) error {
	for {
		select {
		case event := <-w.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				return nil
			}
		case err := <-w.Errors:
			return err
		}
	}
}
