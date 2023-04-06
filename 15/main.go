package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func createHugeString(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString("А") // Тест с 2-байтовыми символами
	}
	return sb.String()
}

func main() {
	someFunc()
	l := len(strings.Split(justString, "")) // Получаем фактическое количество символов
	fmt.Println(l)                          // Не смотря на то, что len равен 100, это 50 символов
	fmt.Println(justString)

	// FIX

	fixed := []rune(createHugeString(1 << 10))
	fixed = fixed[:100]
	fmt.Println(len(strings.Split(string(fixed), "")))
	fmt.Println(string(fixed)) // Теперь он всегда сокращается до постоянного количества символов
}
