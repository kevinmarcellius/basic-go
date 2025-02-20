package logic

import (
	sliceprinter "github.com/kevinmarcellius/go-print-slice"
)

func Hi() string {

	return "Hi, World!"

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

func Logic2_2(n int) {
	matrix := DynamicSlice(n, n)

	value := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][i] = value

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_3(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = value
			value += 2
		}
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_5(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i%2 == 0 {
				matrix[i][j] = value
			} else {
				matrix[i][n-1-j] = value
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

func Logic2_9(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i+j) == (n-1) || i == j {
				matrix[j][i] = value

			}

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_10(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				matrix[j][i] = value

			}

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic2_11(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				matrix[j][i] = value

			}

		}
		value += 2
	}

	sliceprinter.Print2DSlice(matrix)
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
