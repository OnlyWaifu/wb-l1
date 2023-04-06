package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abCksjrAjsChsc"
	b := "aBcDe"
	fmt.Println(allUnique(a))
	fmt.Println(allUnique(b))
}

func allUnique(s string) bool {
	m := make(map[rune]struct{})
	// Диапазон строк в нижнем регистре для проверки без учета регистра
	for _, val := range strings.ToLower(s) {
		// Если map не имеет значения, т.е. значение встречается впервые, то продолжаем ранжирование
		// В противном случае возвращаем false
		if _, ok := m[val]; ok {
			return false
		}
		m[val] = struct{}{}
	}
	return true
}
