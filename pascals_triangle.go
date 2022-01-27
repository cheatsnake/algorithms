package algorithms

import "fmt"

func PascalsTriangle() {
	var deepth int
	rows := [][]int{{1}}

	fmt.Println("Input number:")
	fmt.Scanf("%d\n", &deepth)

	for i := 1; i < deepth; i++ {
		row := []int{}

		for j := 0; j <= len(rows[i-1]); j++ {
			if j == 0 || j == len(rows) {
				row = append(row, 1)
			}
			if j != 0 && j < len(rows[i-1]) {
				row = append(row, (rows[i-1][j-1] + rows[i-1][j]))
			}
		}

		rows = append(rows, row)
	}

	for i := 0; i < len(rows); i++ {
		fmt.Println(rows[i])
	}
}