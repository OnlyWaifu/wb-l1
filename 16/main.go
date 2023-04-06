package main

import "fmt"

func quickSort(arr []int) []int {
	// Если длина слайса меньше 2, то он не нужндается в сортировке
	if len(arr) < 2 {
		return arr
	}

	// Инициализируем переменную pivot, которая является первым элементом слайса
	// Создаём 2 слайса
	pivot := arr[0]
	less := make([]int, 0, 1+len(arr)/2)
	greater := make([]int, 0, 1+len(arr)/2)

	for _, val := range arr[1:] {
		if val <= pivot {
			less = append(less, val) // В слайс less будут добавляться значения, которые меньше или равны pivot
		} else {
			greater = append(greater, val) // А в слайс greater значения, которые больше pivot
		}
	}

	// Добавляем в слайс поочерёдно меньшие или равные значения, pivot, большие значения
	arr = append([]int{}, quickSort(less)...)
	arr = append(arr, pivot)
	arr = append(arr, quickSort(greater)...)
	return arr
}

func main() {
	fmt.Println(quickSort([]int{8, 5, 7, 14, 6, 79, 34, 8}))
}
