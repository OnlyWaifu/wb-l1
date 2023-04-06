package main

import "fmt"

func setIntersetcion(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for _, v := range arr1 {
		m[v] += 1
	}
	for _, v := range arr2 {
		m[v] += 1
	}
	res := make([]int, 0)
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 10}
	set2 := []int{3, 4, 14, 27}
	fmt.Printf("first set: %v\nsecond set: %v\n", set1, set2)

	res := setIntersetcion(set1, set2)
	fmt.Printf("intersection: %v\n", res)
}
