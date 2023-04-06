package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SetName(name string) {
	h.Name = name
}

func (h *Human) SetAge(age int) {
	h.Age = age
}

type Action struct {
	Name       string
	SecondName string
	Age        int
	Human
	H Human
}

func (a *Action) SetName(name string) {
	a.Name = name
}

func (a *Action) SetSecondName(sn string) {
	a.SecondName = sn
}

func (a *Action) SetAge(age int) {
	a.Age = age
}
func main() {
	var action Action = Action{
		Name:       "Sergey",
		SecondName: "Sergeevich",
		Age:        27,
		Human: Human{
			Name: "Evgeniy",
			Age:  19,
		},
		H: Human{
			Name: "Oleg",
			Age:  43,
		},
	}
	// Обращение к полям

	// При наличии одикаковых полей, приоритет у поля более верхнего уровня
	fmt.Println(action.Name) // Sergey

	// Если поле уникальное, то можно напрямую обращаться к нему из родительской структуры
	fmt.Println(action.SecondName) // Sergeevich

	fmt.Println(action.Age) // 27

	// Чтобы обратиться к полю встроенной структуре, нужно указать её тип, затем поле
	fmt.Println(action.Human.Name) // Evgeniy

	// Чтобы обратиться к полю типа структуры, нужно указать имя поля, затем необходимое поле вложенной структуры
	fmt.Println(action.H.Name) // Oleg

	// Вызов методов

	// Методы встроенной структуры наследуются родительской
	// При наличии одинаковых методов, приоритет у метода структуры более верхнего уровня
	action.SetName("Ivan")
	action.SetSecondName("Ivanovich")
	action.SetAge(21)

	// Чтобы вызвать этот метод у встроенной структуры, необходимо обратиться к ней через тип
	action.Human.SetName("Ilya")
	action.Human.SetAge(28)

	fmt.Println(action.Name)       // Ivan
	fmt.Println(action.SecondName) // Ivanovich
	fmt.Println(action.Age)        // 21
	fmt.Println(action.Human.Name) // Ilya
	fmt.Println(action.Human.Age)  // 28
}
