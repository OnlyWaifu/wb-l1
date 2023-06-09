package main

import "fmt"

type bit int

const (
	zero = iota // 0
	one         // 1
)

func main() {
	a := int64(0)
	b := ^a
	fmt.Printf("%d\t%064b\n%d\t%064b\n", a, a, b, b)

	setBit(&a, 10, one)
	setBit(&b, 10, zero)
	fmt.Printf("%d\t%064b\n%d\t%064b\n", a, a, b, b)
}

// Функция setBit устанавливает бит в заданную позицию
func setBit(n *int64, pos int, b bit) {
	switch b {
	case zero:
		*n &= ^(1 << pos)
	case one:
		*n |= 1 << pos
	default:
		panic("invalid bit")
	}
}
