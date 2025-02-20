package main

import (
	"fmt"
	"math"

	sliceprinter "github.com/kevinmarcellius/go-print-slice"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Logic 2.1")
	Logic2_1(10)
	fmt.Println("Logic 2.6")
	Logic2_6(9)
	fmt.Println("Logic 2.7")
	Logic2_7(10)

	fmt.Println("Logic 2.8")
	Logic2_8(10)
	fmt.Println("Logic 2.12")
	Logic2_12(10)

	fmt.Println("Logic 2.13")
	Logic2_13(9)

	fmt.Println("Logic 3.6")
	Logic3_6(9)
}

func DynamicSlice(row, col int) [][]int {

	cell := make([][]int, row)
	for i := range cell {
		cell[i] = make([]int, col)
	}
	return cell
}

func Logic2_12(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			if (i <= j && (i+j) <= (n-1)) || (i >= j && (i+j) >= (n-1)) {
				matrix[j][i] = value
			}

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic1_7(n int) {
	numbers := make([]int, 0)
	fmt.Println(n / 2)
	value := 1
	for i := 0; i < n; i++ {
		if value <= n {
			numbers = append(numbers, value)
		}

		if i < n/2 {
			value += 2
		} else {
			value -= 2
		}
	}

	fmt.Println(numbers)

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

func Logic3_6(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i <= n/2; i++ {

		for j := 0; j < n; j++ {
			if (i <= j && (i+j) <= (n-1)) || (i >= j && (i+j) >= (n-1)) {
				if i%2 == 0 {
					matrix[i][j] = value
					matrix[n-i-1][j] = value
				} else {
					matrix[i][n-1-j] = value
					matrix[n-i-1][n-1-j] = value
				}

				value += 2
			}

		}

	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_1(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i] = value
		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)

}

func Logic2_7(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				matrix[i][j] = value
				value += 2
			}

		}
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_13(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i <= n/2; i++ {
		value = 1
		for j := 0; j < n; j++ {
			if j >= i && (i+j) <= (n-1) {
				matrix[i][j] = value
				matrix[n-1-i][j] = value
			}

			value += 2

		}
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_6(n int) {
	matrix := DynamicSlice(n, n)
	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				matrix[i][j] = value
				value += 3
			} else {
				matrix[i][n-1-j] = value
				value += 2
			}

		}
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_8(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i + j) == (n - 1) {
				matrix[j][i] = value

			}

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
}
