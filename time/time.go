package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// add to date
	add := now.Add(time.Hour * 24 * 5)
	fmt.Println(add)

	// subtract date
	sub := now.Sub(add)
	fmt.Println(sub)

	// formating
	// Type	Placeholder
	// Day	2 or 02 or _2
	// Day of Week	Monday or Mon
	// Month	01 or 1 or Jan or January
	// Year	2006 or 06
	// Hour	03 or 3 or 15
	// Minutes	04 or 4
	// Seconds	05 or 5
	// Milli Seconds  (ms)	.000        //Trailing zero will be includedor .999   //Trailing zero will be omitted
	// Micro Seconds (μs)	.000000             //Trailing zero will be includedor .999999        //Trailing zero will be omitted
	// Nano Seconds (ns)	.000000000        //Trailing zero will be includedor .999999999 //Trailing zero will be omitted
	// am/pm	PM or pm
	// Timezone	MST
	// Timezone offset	Z0700 or Z070000 or Z07 or Z07:00 or Z07:00:00  or -0700 or  -070000 or -07 or -07:00 or -07:00:00
	formated := now.Format("02/01/2006")
	fmt.Println(formated)

	// convert to utc
	loc, _ := time.LoadLocation("UTC")
	utc := time.Now().In(loc)
	fmt.Println(utc)
}
