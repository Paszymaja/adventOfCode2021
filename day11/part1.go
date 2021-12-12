package main

import "fmt"

var data = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func main() {

	var grid [][]int = make([][]int, 10)
	for i := range grid {
		grid[i] = make([]int, len(data[i]))

		for j := 0; j < 10; j++ {
			grid[i][j] = int(data[i][j] - '0')
		}
	}

	fmt.Print(grid)
}
