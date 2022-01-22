package algorithms

import "fmt"

var (
	current = 0
	next = 0
)

func Fibonacci() {
	var (
		pos int
		current = 0
		next = 0
	)

	fmt.Println("Input a number:")
	fmt.Scanf("%d\n", &pos)

	for i := 0; i < pos; i++ {
		switch i {
			case 0: 
				current, next = 0, 1
			case 1: 
				current, next = 1, 1
			case 2:
				current, next = 1, 2
			default:
				result := current + next
				current = next
				next = result
		}
		fmt.Println(current)
	}
}