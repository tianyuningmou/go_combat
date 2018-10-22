/*

Copyright () 2018

All rights reserved

FILE: enumerate.go
AUTHOR: tianyuningmou
EMAIL: tianyuningmou2009@126.com
DATE CREATED:  @Time : 2018/10/11 下午6:00

*/

package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

const (
	FlagNone = 1 << iota
	FlagRed
	FlagBlue
	FlagGreen
)

func main()  {
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)
	fmt.Println(FlagNone, FlagRed, FlagBlue, FlagGreen)
}
