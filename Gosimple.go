package main

import "fmt"

const ss string = "constant string"

func main() {
	var ttt = true
	fmt.Println("ttt value is", ttt)

	var fff int
	fmt.Println(fff)

	fmt.Println("ss value is :", ss)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	if gggt := 1; gggt > 0 {
		fmt.Println("gggt value:", gggt)
	} else {
		fmt.Println("gggt < 0,value is :", gggt)
	}

	var array [5]int
	array[1] = 1111
	fmt.Println("emp :", array)

	slic := make([]string, 4)
	fmt.Println(slic)
	slic[0] = "jjj"
	slic[1] = "ttt"
	slic[2] = "ggg"
	slic[3] = "jjjjj"
	fmt.Print(slic)
	fmt.Println("len: slice", len(slic))
	fmt.Println("slic[3] is", len(slic[3]))

	slic = append(slic, "coconut")
	slic = append(slic, "google", "telsa", "faradayFuture")
	fmt.Println("slic is : ", slic)

	value := make([]string, len(slic))
	copy(value, slic)
	fmt.Println(value)
	fmt.Println("len is ", len(slic))
	split := slic[1:]
	fmt.Println(split, "and len is:", len(split))
	void()
	fmt.Println("lets start map")

	mmap := make(map[string]string)
	mmap["key1"] = "value1"
	mmap["key2"] = "value2"
	fmt.Println("len is", len(mmap))
	delete(mmap, "key1")
	ii, vv := mmap["key2"]
	fmt.Println("ii vv is", ii, vv)
	fmt.Println(mmap)

	nums := []int{2, 3, 4}
	sum := 0
	for index, num := range nums {
		fmt.Println("index :", index)
		sum += num
	}
	fmt.Println("sum :", sum)
	kvs := map[string]string{"aa": "ttt", "gg": "ggg"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

}

func void() {
	fmt.Println("call void method")
}
