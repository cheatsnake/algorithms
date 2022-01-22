package algorithms

import "fmt"

func Euclidean() {
	var (
		current int
		next    int
		suber 	int
	)

	fmt.Println("Input 2 numbers:")
	fmt.Scanf("%d\n", &current)
	fmt.Scanf("%d\n", &next)

	if next > current {
		suber = current
		current = next
		next = suber
	}

	for (current - next) > 0 {
		suber = next
		for current > suber {
			current -= suber
		}
		next = current
		current = suber
	}

	fmt.Println(current)
}