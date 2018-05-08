package main

import (
	"fmt"
	"sync"
	"math"
)

var wg1 sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	functions := make([]func(*sync.WaitGroup), 0)
	functions = append(functions, f1)
	functions = append(functions, f2)
	functions = append(functions, f3)
	for _, f := range functions {
		wg.Add(1)
		go f(&wg)
	}
	wg.Wait()

	r := rect{3, 4}
	c := circle{5}

	structs := make([]shape, 0)
	structs = append(structs, r)
	structs = append(structs, c)

	for _, v := range structs {
		wg1.Add(1)
		go measure(v)
	}
	wg1.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in f1")
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in f2")
}

func f3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in f3")
}

type circle struct {
	r float64
}

type rect struct {
	l float64
	w float64
}

func (r rect) area() float64 {
	defer wg1.Done()
	a := r.l * r.w
	fmt.Println(a)
	return a
}
func (c circle) area() float64 {
	defer wg1.Done()
	a := math.Pi * c.r * c.r
	fmt.Println(a)
	return a

}

type shape interface {
	area() float64
}

func measure(s shape) float64 { return s.area() }
