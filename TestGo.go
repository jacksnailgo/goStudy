package main

import "fmt"

func main() {
	for i, c := range "ab" {
		fmt.Println(i, c)
	}
	fmt.Println(plus(11, 22, 33))
	fmt.Println(ppt("tes", "la", "farday Future"))
	res := sum(111, 3333)
	fmt.Println("res", res)
	fmt.Println(vals(22, 333))
	sums(1, 2)
	sums(2, 2, 2, 2, 2)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sums(nums...)
}

func sum(a int, b int) int {
	return a + b
}

func plus(a, b, c int) int {
	return a + b + c
}

func ppt(a, b, c string) string {
	return a + b + c
}

func vals(a int, b int) (int, int) {
	return a, b
}

func sums(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("sums:", total)
}

/**
多返回值的
*/
