package riddle

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	size := maxValue // Change this to change the size of the matrix
	matrices := GenerateAllMatrices(size)
	fmt.Printf("All possible %dx%d matrices (in 3D slice):\n", size, size)
	for _, matrix := range matrices {
		for _, row := range matrix {
			fmt.Println(row)
		}
		fmt.Println()
	}

}
