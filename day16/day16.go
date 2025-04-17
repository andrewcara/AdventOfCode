package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

type Position struct {
	Y int
	X int
}

type Move struct {
	Position  Position
	Direction int
}

func (p Position) add(other Position) Position {
	return Position{Y: p.Y + other.Y, X: p.X + other.X}
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

func move(current_position Position, grid [][]string, path_lengths map[Position]int) bool {
	// Create a key for the current position
	nodes_list := make([]Move, 0)
	current_move := Move{Position: current_position, Direction: RIGHT}

	nodes_list = append(nodes_list, current_move)

	for len(nodes_list) > 0 {
		current_move = nodes_list[0]

		next_moves := []Move{
			{Position{current_move.Position.Y, current_move.Position.X + 1}, RIGHT},
			{Position{current_move.Position.Y + 1, current_move.Position.X}, DOWN},
			{Position{current_move.Position.Y, current_move.Position.X - 1}, LEFT},
			{Position{current_move.Position.Y - 1, current_move.Position.X}, UP},
		}

		nodes_list = nodes_list[1:]

		for _, move := range next_moves {

			if grid[move.Position.Y][move.Position.X] != "#" {
				if dijsktra(current_move, move, path_lengths) {
					nodes_list = append(nodes_list, move)
				}
			}
		}

	}

	return true
}

func dijsktra(current_move Move, next_move Move, path_lengths map[Position]int) bool {

	move_score := scoreMove(current_move, next_move) + path_lengths[current_move.Position]

	if path_lengths[next_move.Position] > move_score {
		path_lengths[next_move.Position] = move_score
		return true
	}
	return false
}

func scoreMove(current_position Move, next_Position Move) int {
	if current_position.Direction != next_Position.Direction {
		return 1001
	}
	return 1
}

func main() {
	input := getInput()
	grid, start_position, map_values, exit := inputToGrid(input)
	move(start_position, grid, map_values)
	fmt.Println(map_values[exit])
	fmt.Println(map_values[Position{1, 136}])
}

func inputToGrid(input string) ([][]string, Position, map[Position]int, Position) {
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	map_value := make(map[Position]int)
	var start_location Position
	exit := Position{Y: 0, X: 0}
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, char := range line {
			grid[y][x] = string(char)
			position := Position{Y: y, X: x}
			if string(char) == "E" {
				exit = position
			}
			if string(char) == "S" {
				start_location = position
				map_value[position] = 0
			} else {
				map_value[position] = math.MaxInt
			}
		}
	}
	return grid, start_location, map_value, exit
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
