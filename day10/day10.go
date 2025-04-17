package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Y int
	X int
}

func (p Position) isInBounds(row_boundary int, height_boundary int) bool {
	if (p.X < 0 || p.X >= row_boundary) || (p.Y < 0 || p.Y >= height_boundary) {
		return false
	}
	return true
}

func (p Position) isEqual(prev Position) bool {
	if p.Y == prev.Y && p.X == prev.X {
		return true
	}
	return false
}

func main() {
	input := getInput()
	grid := inputToGrid(input)
	num_good_trails := 0
	for row_index, row := range grid {
		for col := range row {
			if grid[row_index][col] == 0 {
				has_visited := make(map[Position]bool)
				loop_trails, _ := move(Position{Y: row_index, X: col}, Position{Y: row_index, X: col}, has_visited, 0, grid)
				num_good_trails += loop_trails
				println(loop_trails)
			}
		}
	}
	println(num_good_trails)
}

func move(current_position Position, prev_position Position, has_visited map[Position]bool, good_trails int, grid [][]int) (int, map[Position]bool) {

	if !current_position.isInBounds(len(grid[0]), len(grid)) {
		has_visited[current_position] = true
		return good_trails, has_visited
	}
	if has_visited[current_position] {
		return good_trails, has_visited
	}
	if grid[current_position.Y][current_position.X]-grid[prev_position.Y][prev_position.X] != 1 && !current_position.isEqual(prev_position) {
		return good_trails, has_visited
	} else if grid[current_position.Y][current_position.X] == 9 && grid[prev_position.Y][prev_position.X] == 8 {
		// has_visited[prev_position] = true
		// has_visited[current_position] = true
		good_trails++
		return good_trails, has_visited
	}
	//println("current postion", current_position.Y, current_position.X, "prev position ", prev_position.Y, prev_position.X)

	moves := []Position{
		{current_position.Y, current_position.X - 1},
		{current_position.Y - 1, current_position.X},
		{current_position.Y, current_position.X + 1},
		{current_position.Y + 1, current_position.X},
	}

	// Recursively explore moves, skipping already visited positions
	for _, moving := range moves {
		good_trails, has_visited = move(moving, current_position, has_visited, good_trails, grid)
	}

	return good_trails, has_visited
}

func inputToGrid(input string) [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for y, line := range lines {
		for x := range line {
			cast_val, _ := strconv.Atoi(string(line[x]))
			grid[y] = append(grid[y], cast_val)
		}
	}
	return grid
}

func getInput() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println(fileName, " file is empty")
		os.Exit(1)
	}
	input := strings.ReplaceAll(string(data), "\r\n", "\n") // doing this replace so it can handle both linux and window text format
	return strings.TrimSpace(input)                         // doing this cause usually there's an extra new line at the bottom of the input
}
