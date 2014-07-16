package main

import “fmt”

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = “Golang”

	fmt.Printf(“Int ist: %d\n”, ms.i1)
	fmt.Printf(“Float ist: %f\n”, ms.f1)
	fmt.Printf(“String ist: %s\n”, ms.str)
	fmt.Println(ms)
}