package main

import "fmt"

const ELEMS = 100000000

type B struct {
	s string
}

type D struct {
	B
	t string
}

func main() {

	var bp [ELEMS]D
	for i := 0; i < ELEMS; i++ {
		bp[i].s = string(int('A') + i%26)
	}
	var x int32
	for i := 0; i < ELEMS; i++ {
		x += int32(bp[i].s[0]) - int32('A')
	}
	fmt.Printf("%d\n", x)
}
