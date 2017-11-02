package main

import "fmt"

type Model struct {

}

func(m Model) get() {
	fmt.Println("woof from model")
}

type Dog struct {
	Model
}

func main() {
	dog := Dog{}

	dog.get()
}
