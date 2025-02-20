package logic

import (
	sliceprinter "github.com/kevinmarcellius/go-print-slice"
)

func Logic4_1(n int) {
	// get matrix size
	size := 0

	for i := 1; i <= n; i++ {
		size += i
	}

	// init matrix
	matrix := DynamicSlice(size, size)
	start := 0
	value := 1
	end := 0
	for i := 1; i <= n; i++ {
		end = i + start - 1
		for j := 0; j < i; j++ {

			for k := 0; k < i; k++ {

				if j%2 != 0 {
					matrix[j+start][end-k] = value
				} else {
					matrix[j+start][k+start] = value
				}
				value += 2
			}

		}
		value = 1
		start += i
	}

	sliceprinter.Print2DSlice(matrix)
}

