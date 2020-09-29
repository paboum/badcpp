package main

import "fmt"

type S struct {
	s string
}

type PS struct {
	ps *string
}

func main() {

	var s S = S{"A"}
	var t S = S(s)
	s.s = "B"
	fmt.Printf("%s\n", s.s)
	fmt.Printf("%s\n", t.s)

	t = s
	t.s = "C"
	fmt.Printf("%s\n", s.s)
	fmt.Printf("%s\n", t.s)

	var x PS = PS{&s.s}
	var y PS = PS(x)
	*x.ps = "D"
	fmt.Printf("%s\n", s.s)
	*y.ps = "E"
	fmt.Printf("%s\n", s.s)
}
