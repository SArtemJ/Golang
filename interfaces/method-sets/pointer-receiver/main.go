package main

import (
	"math"
	"fmt"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

//func (c *circle) area() float64 { - work fine with &circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main()  {
	c := circle{10}
	info(&c)
}