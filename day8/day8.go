package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	Y, X int
}

func (p Position) isInBounds(row_boundary int, height_boundary int) bool {
	if (p.X < 0 || p.X >= row_boundary) || (p.Y < 0 || p.Y >= height_boundary) {
		return false
	}
	return true
}

func (p Position) Add(other Position) Position {
	return Position{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}
func (p Position) isEqual(curr_p Position) bool {
	if p.X == curr_p.X && p.Y == curr_p.Y {
		return true
	}
	return false
}

func main() {

	input := getInput()
	part1(input)
}

func part1(input string) {
	grid := inputToGrid(input)
	anti_node_map := make(map[Position]bool)

	for Y, row := range grid {

		for X, row_val := range row {
			if row_val != "." {
				visited := make(map[Position]bool)
				startPos := Position{Y: Y, X: X}
				visited, anti_node_map = move(startPos, visited, anti_node_map, startPos, grid)
			}
		}
	}
	fmt.Println(len(anti_node_map))
}

func move(current_position Position, visited map[Position]bool, anti_node_map map[Position]bool, start_position Position, grid [][]string) (map[Position]bool, map[Position]bool) {
	if !current_position.isInBounds(len(grid[0]), len(grid)) {
		return visited, anti_node_map
	}

	if visited[current_position] {
		return visited, anti_node_map
	}
	visited[current_position] = true

	if string(grid[current_position.Y][current_position.X]) == string(grid[start_position.Y][start_position.X]) &&
		!current_position.isEqual(start_position) {
		anti_node_map[current_position] = true
		anti_node_map[start_position] = true
		anti_node_val := add_anti_node(current_position, start_position)
		new_start := current_position

		for anti_node_val.isInBounds(len(grid[0]), len(grid)) {
			anti_node_map[anti_node_val] = true

			temp := Position{Y: anti_node_val.Y, X: anti_node_val.X}
			anti_node_val = add_anti_node(anti_node_val, new_start)
			new_start = temp

		}
		return visited, anti_node_map
	}

	// Define possible moves
	moves := []Position{
		{current_position.Y, current_position.X - 1},
		{current_position.Y - 1, current_position.X},
		{current_position.Y, current_position.X + 1},
		{current_position.Y + 1, current_position.X},
	}

	// Recursively explore moves, skipping already visited positions
	for _, moving := range moves {
		visited, anti_node_map = move(moving, visited, anti_node_map, start_position, grid)
	}

	return visited, anti_node_map
}

func add_anti_node(current_position, start_position Position) Position {
	return Position{
		Y: current_position.Y + (current_position.Y - start_position.Y),
		X: current_position.X + (current_position.X - start_position.X),
	}
}

func inputToGrid(input string) [][]string {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))

	for y, line := range lines {
		for x := range line {
			grid[y] = append(grid[y], string(line[x]))

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
