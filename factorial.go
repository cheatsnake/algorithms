package algorithms

import "fmt"

func Factorial() {
	var result int = 1
	var num int

	fmt.Println("Input a number:")
	fmt.Scanf("%d\n", &num)

	for i := 1; i <= num; i++ {
		result *= i
	}

	fmt.Println(result)
}