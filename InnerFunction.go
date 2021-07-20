package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 22})

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
