package main

import "fmt"

func main() {
	a := []int{10, 23, 35, 46, 57, 68, 72, 81, 94}
	r := binarySearchRecursive(a, 57)
	i := binarySearchIterative(a, 57)
	fmt.Printf("recursive: %d\niterative: %d\n", r, i)
}

// Рекурсивная реализация
func binarySearchRecursive[T ~int](a []T, val T) int {
	mid := len(a) / 2 // middle index

	// Проверка равенства
	if a[mid] == val {
		return mid
	}

	// Рекурсивный вызов в правой половине
	if a[mid] < val {
		return binarySearchRecursive(a[mid:], val)
	}

	// Рекурсивный вызов в левой половине
	if a[mid] > val {
		return binarySearchRecursive(a[:mid], val)
	}

	// Если значение не существует в массиве
	return -1
}

func binarySearchIterative[T ~int](a []T, val T) int {
	low, high := 0, len(a)-1 // Индексы для крайнего левого и крайнего правого значения
	mid := high / 2          // Индекс среднего элемента

	// Начало поиска
	for low <= high {
		// Проверка на равенство
		if a[mid] == val {
			return mid
		}

		// Изменение левой и средней части на поиск в правой половине
		if a[mid] < val {
			low = mid + 1
			mid = (high - low) / 2
			continue
		}

		// Изменение середней и правой части на поиск в левой половине
		high = mid - 1
		mid = (high - low) / 2
	}

	// Если значение не существует в массиве
	return -1
}
