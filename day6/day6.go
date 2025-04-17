package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type guard struct {
	x          int
	y          int
	dirIndex   int
	directions [4][2]int
}

func (G *guard) getNextPos() (int, int) {
	return G.x + G.directions[G.dirIndex%4][0], G.y + G.directions[G.dirIndex%4][1]
}

func (G *guard) turnRight() {
	G.dirIndex += 1
}

func (G *guard) stepForward() {
	G.x, G.y = G.getNextPos()
}

func main() {
	input := getInput()
	t := time.Now()
	fmt.Println("Part 1 answer:", part1(input), " (in", time.Since(t).Milliseconds(), "ms)")
	t = time.Now()
	fmt.Println("Part 2 answer:", part2(input), " (in", time.Since(t).Milliseconds(), "ms)")
}

// guard walks around and leaves X on the ground, then we count all X's left on the ground
func part1(input string) string {
	grid, theGuard := inputToGrid(input)
	grid[theGuard.y][theGuard.x] = 'X'
	for isInBounds(theGuard.x, theGuard.y, grid) {
		newX, newY := theGuard.getNextPos()
		if !isInBounds(newX, newY, grid) {
			theGuard.stepForward()
		} else if grid[newY][newX] == '#' {
			theGuard.turnRight()
		} else {
			theGuard.stepForward()
			grid[theGuard.y][theGuard.x] = 'X'
		}
	}

	countX := 0
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == 'X' {
				countX++
			}
		}
	}

	return fmt.Sprint(countX)
}

// We make the guard walk on her path, but on each stepForward we temporarily put a wall in her face,
// and then we create a clone guard that will walk from the guard's original spot and we see if the new wall
// will make the clone loop around
// if it does loop, leave an X on the ground where the real guard is standing at, and then count all the X's
func part2(input string) string {
	grid, theGuardOG := inputToGrid(input)
	theGuard := theGuardOG
	gridCopy := make([][]rune, len(grid))
	for i := 0; i < len(gridCopy); i++ {
		gridCopy[i] = slices.Clone(grid[i])
	}

	for isInBounds(theGuard.x, theGuard.y, grid) {
		newX, newY := theGuard.getNextPos()
		if !isInBounds(newX, newY, grid) {
			theGuard.stepForward()
		} else if grid[newY][newX] == '#' {
			theGuard.turnRight()
		} else {
			grid[newY][newX] = '#'
			if gridCopy[newY][newX] != 'X' && stateLoops(theGuardOG, grid) {
				gridCopy[newY][newX] = 'X'
			}
			grid[newY][newX] = '.'
			theGuard.stepForward()
		}
	}

	countX := 0
	for y := range gridCopy {
		for x := range gridCopy[0] {
			if gridCopy[y][x] == 'X' {
				countX++
			}
		}
	}

	return fmt.Sprint(countX)
}

// find a loop by remembering each time the guard collides with a wall, we save the wall's coordinates and the direciton
// the guard was walking towards when it collided with that wall. If we see this combination again we know it's in a loop
func stateLoops(theGuard guard, grid [][]rune) bool {
	guardClone := theGuard
	walls := [][3]int{} // 3 numbers ->wall's coordinates: x, y, direction (0-3) < guard's direction

	for isInBounds(guardClone.x, guardClone.y, grid) {
		newX, newY := guardClone.getNextPos()
		if !isInBounds(newX, newY, grid) {
			guardClone.stepForward()
		} else if grid[newY][newX] == '#' {
			wall := [3]int{newX, newY, guardClone.dirIndex % 4}
			if slices.Contains(walls, wall) {
				return true
			} else {
				walls = append(walls, wall)
			}
			guardClone.turnRight()
		} else {
			guardClone.stepForward()
		}
	}
	return false
}

func inputToGrid(input string) ([][]rune, guard) {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	theGuard := guard{}
	theGuard.directions = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	for y, line := range lines {
		for x, tile := range line {
			switch tile {
			case '^':
				theGuard.x = x
				theGuard.y = y
				grid[y] = append(grid[y], '.')
			default:
				grid[y] = append(grid[y], tile)
			}
		}
	}
	return grid, theGuard
}

func isInBounds(x, y int, grid [][]rune) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
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
