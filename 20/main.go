package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}

func reverseWords(s string) string {
	ss := strings.Split(s, " ") // Разделение на слова

	// Меняем местами i-е и j-е значения, пока j >= i
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}

	// Возвращаем строку с перевернутыми словами
	return strings.Join(ss, " ")
}
