package main

import (
	"fmt"
	"os"
)

func main() {
	number := input()
	fmt.Printf("factorialIterative = %d, factorialRecursive = %d", getIterativeFactorial(number),
		getRecursiveFactorial(number, 1))
}

func input() int {
	fmt.Printf("Введите число для проверки: ")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		os.Exit(404)
	}
	return number
}

func getIterativeFactorial(step int) int {
	var result = 1
	for step != 1 {
		result *= step
		step--
	}
	return result
}

func getRecursiveFactorial(step int, result int) int {
	if(step == 1) {
		return result
	}
	return getRecursiveFactorial(step-1,result*step)
}