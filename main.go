package main

import "fmt"

func main() {
	helloWorld()
	age()
}

func age() {
	const name, age = "Abhishek", 23
	fmt.Println(name, " is ", age, " years old!")
}

func helloWorld() {
	fmt.Println("Hello world!")
}
