package riddle

func getTaskInit() [][]Cell {
	matrix := [][]Cell{
		{NewCell(1), NewCell(1), NewCell(4), NewCell(3), NewCell(4), NewCell(1), NewCell(3), NewCell(2), NewCell(2)},
		{NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(2), NewCell(1), NewCell(3), NewCell(2), NewCell(2)},
		{NewCell(3), NewCell(2), NewCell(1), NewCell(4), NewCell(3), NewCell(3), NewCell(2), NewCell(1), NewCell(3)},
		{NewCell(4), NewCell(3), NewCell(4), NewCell(2), NewCell(3), NewCell(1), NewCell(1), NewCell(2), NewCell(4)},
		{NewCell(4), NewCell(2), NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(3), NewCell(4), NewCell(1)},
		{NewCell(2), NewCell(2), NewCell(3), NewCell(3), NewCell(4), NewCell(4), NewCell(4), NewCell(1), NewCell(2)},
		{NewCell(2), NewCell(3), NewCell(3), NewCell(1), NewCell(3), NewCell(2), NewCell(2), NewCell(4), NewCell(1)},
		{NewCell(4), NewCell(4), NewCell(2), NewCell(1), NewCell(3), NewCell(1), NewCell(2), NewCell(3), NewCell(3)},
		{NewCell(4), NewCell(4), NewCell(2), NewCell(1), NewCell(1), NewCell(1), NewCell(2), NewCell(3), NewCell(3)},
	}
	return matrix
}

func getExampleInit() [][]Cell {
	matrix := [][]Cell{
		{NewCell(4), NewCell(2), NewCell(4), NewCell(8)},
		{NewCell(8), NewCell(6), NewCell(6), NewCell(8)},
		{NewCell(4), NewCell(2), NewCell(6), NewCell(6)},
		{NewCell(2), NewCell(2), NewCell(6), NewCell(6)},
	}
	return matrix
}

func GetExampleResult() [][]Cell {
	matrix := [][]Cell{
		{Cell{true, 4}, Cell{false, 2}, Cell{false, 4}, Cell{false, 8}},
		{Cell{false, 8}, Cell{false, 6}, Cell{true, 6}, Cell{true, 8}},
		{Cell{false, 4}, Cell{true, 2}, Cell{true, 6}, Cell{true, 6}},
		{Cell{false, 2}, Cell{true, 2}, Cell{true, 6}, Cell{true, 6}},
	}
	return matrix
}
