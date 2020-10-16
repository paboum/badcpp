package main

import (
	"fmt"
)

type Vehicle interface {
	weight() float64
	start()
}

type RoadVehicle struct {
	w     float64
	state bool
}

func (r RoadVehicle) weight() float64 {
	return r.w
}

func (r RoadVehicle) weight2() float64 {
	return r.w
}

func (r RoadVehicle) start() {
	r.state = true
}

type Automobile struct {
	RoadVehicle
	Wheels int8
}

type Aircraft struct {
	engines int8
}

/*func (r Aircraft) weight() float64 {
	return 1
}

func (r Aircraft) start() {
}*/

type Helicopter struct {
	Aircraft
	Color int
}

func main() {
	var x [10]Vehicle
	x[0] = new(RoadVehicle)
	var a Automobile
	a.Wheels = 16
	x[1] = &a
	//x[2] = new(Aircraft)
	//x[3] = new(Helicopter)
	/*var v = new(Vehicle)
	x[4] = *v
	x[4].start()
	fmt.Println(x[4])*/
	fmt.Println(x[0].weight())
	//fmt.Println(x[1].Wheels)
	fmt.Println(x[1].(*Automobile).weight2())
	//fmt.Println(x[1].(*Helicopter).Color)
	x[1], x[2] = x[2], x[1]
	fmt.Println(x[1])
	fmt.Println(x[2])
}
