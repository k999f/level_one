package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayName() string {
	return "my name is " + h.Name
}

type Action struct {
	Human
	Greeting string
}

func (a *Action) IntroduceYourself() {
	fmt.Printf("%s, %s, I am %d years old", a.Greeting, a.SayName(), a.Age)
}

func main() {
	// Создаем структуру Human
	somebody := Human{
		Name: "Ivan Ivanov",
		Age:  25,
	}
	// Cоздаем структуру Action с вло;енной структурой Human
	action := Action{
		Human:    somebody,
		Greeting: "Hello",
	}
	// Вызываем из структуры Action метод, который использует метод структуры Human
	action.IntroduceYourself()
}
