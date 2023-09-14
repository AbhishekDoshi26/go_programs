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
	simpleTypes()
	average()
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

func simpleTypes() {
	// simplest way to declare single variable
	var c int = 10
	fmt.Println(c)

	// only inside functions this is allowed. It is called short declaration operator
	d := 10.3
	// show type of variable
	fmt.Printf("The type of d is %T and value is %v\n", d, d)

	// This means that %v needs to use the value of 1st parameter. 0th parameter is the string
	// that we have passed in Printf
	fmt.Printf("The type of d is %T and value is %[1]v\n", d)

	// way of declaring multiple variables
	var (
		a int
		b float64
	)
	a = 10
	b = 2.9
	fmt.Println(a, b)

	// Typecasting a variable
	var e int = 10
	var f float64 = 10.8
	fmt.Println(e, f)

	e = int(f)
	fmt.Println(e)
}

func average() {
	var sum float64
	var n int

	fmt.Println("Enter numbers:")

	for {
		var val float64

		// Fscanln returns 2 values, result of Fscanln and error
		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			break
		}

		sum += val
		n++
	}
	if n == 0 {
		fmt.Fprintln(os.Stderr, "No numbers were found")
		os.Exit(1)
	}

	fmt.Println("The average value is: ", sum/float64(n))

	// If we have a file that contains numbers and if we run go run main.go < num.txt
	// It takes all the values from file in the standard input (stdin) and would calculate the average.
}
