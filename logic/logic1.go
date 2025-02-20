package logic

import (
	"fmt"
	"math"

	sliceprinter "github.com/kevinmarcellius/go-print-slice"
)

func DynamicSlice(row, col int) [][]int {

	cell := make([][]int, row)
	for i := range cell {
		cell[i] = make([]int, col)
	}
	return cell
}

func Logic1_1(n int) {
	cell := make([]int, n)
	value := 1
	for i := 0; i < n; i++ {
		cell[i] = value
		value += 2
	}
	sliceprinter.PrintSlice(cell)
}

func Logic1_2(n int) {
	cell := make([]int, n)
	value := 2
	for i := 0; i < n; i++ {
		cell[i] = value
		value += 2
	}
	sliceprinter.PrintSlice(cell)
}

func Logic1_3(n int) {
	cell := make([]int, n)
	value := 3
	for i := 0; i < n; i++ {
		cell[i] = value
		value += 3
	}
	sliceprinter.PrintSlice(cell)
}

func Logic1_4(n int) {
	cell := make([]int, n)
	value := 2 * n - 1

	for i := 0; i < n; i++ {
		cell[i] = value
		value -= 2
	}

	sliceprinter.PrintSlice(cell)
}

func Logic1_5(n int) {
	cell := make([]int, n)
	value := 2 * n
	for i := 0; i < n; i++ {
		cell[i] = value
		value -= 2
	}
	sliceprinter.PrintSlice(cell)
}
func Logic1_10(n int) {
	value := 1
	for i := 0; i < n; i++ {
		if (i % 2) == 0 {
			fmt.Printf("%d ", int(math.Pow(2, float64(value))))
			value++
		} else {
			fmt.Printf("fizz ")
		}
	}
	fmt.Println()
}

func Logic1_11(n int) {
	start := 1
	value := 3
	for i := 0; i < n; i++ {
		if (i % 2) == 0 {
			fmt.Printf("buzz ")
		} else {
			if i == 1 {
				fmt.Printf("%d ", start)

			} else {
				fmt.Printf("%d ", value)
				value *= 2
			}
		}
	}
	fmt.Println()
}
