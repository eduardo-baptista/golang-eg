package main

import "runtime/debug"

// ultimo que entra primeiro que sai
func f3() {
	debug.PrintStack()
}

func f2() {
	f3()
}

func f1() {
	f2()
}

func main() {
	f1()
}
