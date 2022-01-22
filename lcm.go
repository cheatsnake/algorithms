package algorithms

import (
	"fmt"
	"math"
)

// LCM - Least common multiple
func Lcm() {
	var (
		current int
		next    int
		mult    int
		suber 	int
	)

	fmt.Println("Input 2 numbers:")
	fmt.Scanf("%d\n", &current)
	fmt.Scanf("%d\n", &next)

	mult = int(math.Abs(float64(current * next)))

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

	fmt.Println(mult / current)
}