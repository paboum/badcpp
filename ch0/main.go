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

	t2 := encapsulated.TraceCustom(os.Stderr)
	t2.On()
	t2.Print("X2")
	t2.Off()
	t2.Print("Y2")

	var t3 encapsulated.Trace
	t3 = encapsulated.TraceDefault()
	t3.On()
	t3.Print("X3")
	t3.Off()
	t3.Print("Y3")
	t3 = t2
	t3.On()
	t3.Print("Z3")
}
