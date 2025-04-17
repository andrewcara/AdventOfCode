package main

import (
	"fmt"
	"os"
	"strings"
)

const VISITED = "1"

const (
	RIGHT = 0
	DOWN  = 1
	LEFT  = 2
	UP    = 3
)

type Position struct {
	Y      int
	X      int
	Letter string
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

func main() {
	part1()
	//part2()
}

func part1() {
	input := getInput()
	grid := inputToGrid(input)
	output := 0
	for Y, row := range grid {
		for X, letter := range row {
			if letter != VISITED {
				perimeter_map := make(map[Position]map[int]bool)
				perimeter, area, _ := move(letter, Position{Y: Y, X: X}, &grid, 0, 0, perimeter_map)
				output += (perimeter * area)
				println(perimeter, area)
			}
		}
	}
	fmt.Println(output)
}

func move(start_letter string, current_position Position, grid *[][]string, perimeter int, area int, perimeter_map map[Position]map[int]bool) (int, int, map[Position]map[int]bool) {

	if !current_position.isInBounds(len((*grid)[0]), len(*grid)) {
		return perimeter, area, perimeter_map
	}
	current_position.Letter = (*grid)[current_position.Y][current_position.X]
	if current_position.Letter == strings.ToLower(current_position.Letter) || current_position.Letter != start_letter {
		return perimeter, area, perimeter_map
	} else {
		temp := 0
		temp, perimeter_map = check_perimeter(current_position, *grid, perimeter_map)
		perimeter += temp
		area++
		(*grid)[current_position.Y][current_position.X] = (strings.ToLower(current_position.Letter))
	}

	next_moves := []Position{
		{current_position.Y, current_position.X - 1, ""},
		{current_position.Y - 1, current_position.X, ""},
		{current_position.Y, current_position.X + 1, ""},
		{current_position.Y + 1, current_position.X, ""},
	}

	for _, moves := range next_moves {
		perimeter, area, perimeter_map = move(start_letter, moves, grid, perimeter, area, perimeter_map)
	}
	//fmt.Println(perimeter_map)
	return perimeter, area, perimeter_map
}

//finding the sides of a shape that is represented by a
//traverse 2d slice with the given rule
//1) direction change adds a side if the square is unvisited
//2)cycle of moves is right->down->left->up
//3) we are always checking if we can perform the previous move when we are going in a direction
// i.e. if we are going right we are always checking for a avaialble space that is up
//4) when a square has been left mark it as so. this is important as we can be in a 1x1 square and need to try every direction
// to figure out sides

func getMoveIndex(index int, moves []Position) int {
	return (index + len(moves)) % len(moves)
}

func check_perimeter(p Position, grid [][]string, perimeter_map map[Position]map[int]bool) (int, map[Position]map[int]bool) {

	//actually counting sides here but lazy
	perimeter := 0
	perimeter_map[p] = make(map[int]bool)
	next_moves := []Position{
		{p.Y, p.X - 1, p.Letter},
		{p.Y - 1, p.X, p.Letter},
		{p.Y, p.X + 1, p.Letter},
		{p.Y + 1, p.X, p.Letter},
	}

	for i, move := range next_moves {
		adjacent_index1 := getMoveIndex(i-1, next_moves)
		adjacent_index2 := getMoveIndex(i+1, next_moves)

		if !move.isInBounds(len(grid), len(grid[0])) {
			//println(p.Y, p.X, next_moves[adjacent_index1].Y, next_moves[adjacent_index1].X, next_moves[adjacent_index2].Y, next_moves[adjacent_index2].X)
			//fmt.Println(perimeter_map)
			perimeter_map[p][i] = true
			if !perimeter_map[(next_moves[adjacent_index1])][i] && !perimeter_map[(next_moves[adjacent_index2])][i] {
				println(p.Y, p.X, p.Letter, move.Letter, next_moves[adjacent_index1].Y, next_moves[adjacent_index1].X, next_moves[adjacent_index2].Y, next_moves[adjacent_index2].X, "index", i)
				perimeter++
			}

		} else {
			move.Letter = grid[move.Y][move.X]
			if move.Letter != p.Letter && move.Letter != strings.ToLower(p.Letter) {
				perimeter_map[p][i] = true
				if !perimeter_map[(next_moves[adjacent_index1])][i] && !perimeter_map[(next_moves[adjacent_index2])][i] {

					println(p.Y, p.X, p.Letter, move.Letter, next_moves[adjacent_index1].Y, next_moves[adjacent_index1].X, next_moves[adjacent_index2].Y, next_moves[adjacent_index2].X, "index", i)
					perimeter++
				}
			}
		}
	}
	//fmt.Print(perimeter_map)
	//println("perimeter", perimeter, "coordinates", p.Y, p.X, "Letter", p.Letter)
	return perimeter, perimeter_map

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
