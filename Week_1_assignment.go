package main

import (
	"fmt"
)

func main() {
	var a, b int
	var opr string
	fmt.Scanf("%d %s %d", &a, &opr, &b)
	switch opr {
	case "+":
		fmt.Println(a, " + ", b, " = ", a+b)
	case "-":
		fmt.Println(a, " - ", b, " = ", a-b)
	case "*":
		fmt.Println(a, " * ", b, " = ", a*b)
	case "/":
		if b != 0 {
			fmt.Println(a, " / ", b, " = ", a/b)
		} else {
			fmt.Println("err: Divide by zero")
		}
	default:
		fmt.Println("Irrecognizable operation")
	}
}
