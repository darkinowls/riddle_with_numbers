package riddle

type Cell struct {
	IsMarked bool `json:"IsMarked"`
	Value    int8 `json:"Value"`
}

func NewCell(value int8) Cell {
	return Cell{
		IsMarked: true,
		Value:    value,
	}
}

func TranslateToCells(input [][]int) [][]Cell {
	numRows := len(input)
	numCols := len(input[0])

	output := make([][]Cell, numRows)
	for i := 0; i < numRows; i++ {
		output[i] = make([]Cell, numCols)
		for j := 0; j < numCols; j++ {
			output[i][j] = Cell{
				Value:    int8(input[i][j]),
				IsMarked: true,
			}
		}
	}
	return output
}
