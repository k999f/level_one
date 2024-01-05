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
	somebody := Human{
		Name: "Ivan Ivanov",
		Age:  25,
	}
	action := Action{
		Human:    somebody,
		Greeting: "Hello",
	}
	action.IntroduceYourself()
}
