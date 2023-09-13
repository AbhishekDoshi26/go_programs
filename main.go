package main

import (
	"fmt"
	"os"

	"./hello"
)

func main() {
	helloWorld()
	age()
	useOfOs()
}

func age() {
	const name, age = "Abhishek", 23
	fmt.Println(name, "is", age, "years old!")
}

func helloWorld() {
	fmt.Println(hello.Say([]string{"World"}))
}

func useOfOs() {
	fmt.Println(hello.Say(os.Args[1:]))
}
