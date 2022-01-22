package algorithms

import "fmt"

func SieveOfEratosthenes() {
	var (
		max int
		numbers []int
		primes []interface{}
		prime = 2
	)

	fmt.Println("Input a number:")
	fmt.Scanf("%d\n", &max)

	for i := 2; i <= max; i++ {
		numbers = append(numbers, i)
	}

	for _, v := range numbers {
		if v != prime && v != 0 {
			prime = v
		}
		for i, num := range numbers {
			if num != prime && num % prime == 0 {
				numbers[i] = 0
			}
		} 
	}

	for _, v := range numbers {
		if v != 0 { primes = append(primes, v) }
	}

	fmt.Println(primes...)
}