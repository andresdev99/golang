package utils

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	result := operacionesMiddleware(sumar)(1, 2)
	fmt.Println(result)
	result = operacionesMiddleware(restar)(10, 5)
	fmt.Println(result)
	result = operacionesMiddleware(multiplicar)(6, 7)
	fmt.Println(result)

}

func operacionesMiddleware(functionName func(int, int) int) func(c, d int) int {
	return func(c, d int) int {
		return functionName(c, d)
	}
}
