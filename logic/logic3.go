package logic

import (
	sliceprinter "github.com/kevinmarcellius/go-print-slice"
)

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

func Logic3_7(n int) {
	matrix := DynamicSlice(n, n)

	value := 1
	add := 0
	for i := 0; i <= n/2; i++ {
		for j := 0; j < n; j++ {
			if ((i + j) >= (n / 2)) && ((i+j) <= (n/2 + add)) {
				if i%2 != 0 {
					matrix[i][j] = value
					matrix[n-i-1][j] = value
				} else {
					matrix[i][n-1-j] = value
					matrix[n-i-1][n-1-j] = value
				}
				value += 2
			}

		}
		add += 2
	}

	sliceprinter.Print2DSlice(matrix)
}

func Logic3_14(n int) {
	matrix := DynamicSlice(n, n)

	value := 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i%2 != 0 {
				matrix[n-1-j][i] = value
			} else {
				matrix[j][i] = value
			}
			value += 2
		}
	}

	sliceprinter.Print2DSlice(matrix)
}
