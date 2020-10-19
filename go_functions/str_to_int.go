package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var str string
	total, power := 0.0, 0.0
	fmt.Scan(&str)
	for i := len(str) - 1; i >= 0; i-- {
		num, err := strconv.ParseFloat(str[i: i + 1], 64)
		if 0 <= num && num <= 9 && err == nil{
			total += math.Pow(10, power) * num
			power += 1
		}
	}
	fmt.Print(total)
}
