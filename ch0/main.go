package main

import (
	"os"

	"../ch0/encapsulated"
	"../ch0/public"
)

func main() {
	t1 := public.Trace{F: os.Stdout}
	t1.On()
	t1.Print("X1")
	t1.Off()
	t1.Print("Y1")

	//t2 := encapsulated.TraceCustom(os.Stderr)
	t2 := encapsulated.TraceDefault()
	t2.On()
	t2.Print("X2")
	t2.Off()
	t2.Print("Y2")
}
