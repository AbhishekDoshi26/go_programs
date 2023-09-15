package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./hello"
)

func main() {
	helloWorld()
	age()
	useOfOs()
	simpleTypes()
	//average()
	stringExamples()
	searchAndReplace()
}

func age() {
	const name, age = "Abhishek", 23
	fmt.Println(name, "is", age, "years old! \n\n")
}

func helloWorld() {
	fmt.Println(hello.Say([]string{"World"}) + "\n\n")
}

func useOfOs() {
	fmt.Println(hello.Say(os.Args[1:]) + "\n\n")
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
	fmt.Println("\n\n")
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

	fmt.Println("The average value is: ", sum/float64(n), "\n\n")

	// If we have a file that contains numbers and if we run go run main.go < num.txt
	// It takes all the values from file in the standard input (stdin) and would calculate the average.
}

func stringExamples() {
	s := "èlite"
	fmt.Printf("Type of s is %T and value is %[1]v\n", s)
	// rune means character in int32 form. So basically a string is set of runes.
	// In the output, the first value is 232 (>127), because è is not a unicode character.
	fmt.Printf("Type of s is %T and value is %[1]v\n", []rune(s))

	// Hence, when the 1st rune is converted to byte, it gets expanded
	// and the below will have 6 entries instead of 5
	fmt.Printf("Type of s is %T and value is %[1]v\n", []byte(s))

	// And because a non-unicode character occupies 2 bytes,
	// the length of string is 6 instead of 5!!!!!!!!!!!!!!!!!!!!!!!!
	// This proves that length of a string in the program is
	// number of bytes it needs to store that string.
	fmt.Printf("Length of string is: %d\n\n", len(s))
}

func searchAndReplace() {
	// run with command line.
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Less arguments found!")
		os.Exit(-1)
	}

	// old will be replaced by new
	old, new := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		// split original string from the old words (separator)
		originalString := strings.Split(scan.Text(), old)

		// add new words to the string and join it.
		replacedString := strings.Join(originalString, new)
		fmt.Println(replacedString)
	}
	fmt.Println("\n\n")

}
