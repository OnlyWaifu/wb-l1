package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	i := 2                        // Индекс элемента для удаления
	a = append(a[:i], a[i+1:]...) // Удаление
	fmt.Println(a)
}
