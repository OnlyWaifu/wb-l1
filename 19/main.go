package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}

func reverse(s string) string {
	rs := []rune(s) // Преобразование строки в слайс рун

	// Меняем местами i-е и j-е значения, пока j >= i
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	// Обратное преобразование в строку
	return string(rs)
}
