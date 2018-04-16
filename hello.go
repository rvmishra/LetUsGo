package main

import "fmt"
import "time"
import "math/rand"

func main() {
	fmt.Println("hello, world")
	fmt.Println("time is ", time.Now())
	fmt.Println("random number is ", rand.Intn(100))
	fmt.Println(add(1, 2))
	s1 := "hello"
	s2 := "go"
	s3 := "world"
	fmt.Println(swap(s1, s2, s3))
	java, golang, ruby := 1, 2, 3
	fmt.Println(java, golang, ruby)
	fmt.Println(add1(1, 2))
	fmt.Println(Vertex{1, 2})
	p := Vertex{1, 2}
	swapStruct(&p)
	fmt.Println(p)
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:3])
}

func add1(i, j int) int {
	return i + j
}

func swap(s1, s2, s3 string) (string, string, string) {
	return s3, s2, s1
}

func add(a, b int) int {
	c := a + b
	return c
}

type Vertex struct {
	X int
	Y int
}

func swapStruct(v *Vertex) {
	v = &Vertex{3,4}
}


