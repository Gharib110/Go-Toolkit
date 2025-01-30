package main

import "Go-Toolkit/toolkit"

func main() {
	t := &toolkit.Tools{}

	s := t.RandomString(10)
	println(s)
}
