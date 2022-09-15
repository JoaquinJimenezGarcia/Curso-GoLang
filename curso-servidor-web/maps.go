package main

import "fmt"

func main() {
	map1 := make(map[string]int)

	map1["a"] = 8
	map1["b"] = 10

	fmt.Println(map1)
	fmt.Println(map1["b"])
}
