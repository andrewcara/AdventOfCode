package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Result struct {
	AX, AY, BX, BY, PRIZE_X, PRIZE_Y int
}

func main() {
	input := getInput()
	results := parseInput(input)
	output := 0
	for _, val := range results {
		output += cramer(val, 10000000000000)
		fmt.Println(val)
	}
	println(output)
}

func cramer(machine Result, offset int) int {
	// Adjust the prize values with the offset
	prizeX := machine.PRIZE_X + offset
	prizeY := machine.PRIZE_Y + offset

	// Calculate determinant
	det := machine.AX*machine.BY - machine.AY*machine.BX
	if det == 0 {
		// No solution or infinite solutions
		return 0
	}

	// Solve for A and B using Cramer's rule
	a := (prizeX*machine.BY - prizeY*machine.BX) / det
	b := (machine.AX*prizeY - machine.AY*prizeX) / det
	//println(a, b)
	// Verify the solution
	if machine.AX*a+machine.BX*b == prizeX && machine.AY*a+machine.BY*b == prizeY {
		return a*3 + b
	}

	// Invalid solution
	return 0
}

func parseInput(input string) []Result {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var results []Result
	//fmt.Println(lines)
	for i := 0; i < len(lines); i += 4 {
		// Parse Prize line
		prizeLine := strings.Split(strings.Split(lines[i+2], ": ")[1], ", ")
		prizeX, _ := strconv.Atoi(strings.Split(prizeLine[0], "=")[1])
		prizeY, _ := strconv.Atoi(strings.Split(prizeLine[1], "=")[1])

		// Parse Button A line
		buttonALine := strings.Split(strings.Split(lines[i], ": ")[1], ", ")
		axOffset, _ := strconv.Atoi(strings.TrimPrefix(buttonALine[0], "X+"))
		ayOffset, _ := strconv.Atoi(strings.TrimPrefix(buttonALine[1], "Y+"))

		// Parse Button B line
		buttonBLine := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")
		bxOffset, _ := strconv.Atoi(strings.TrimPrefix(buttonBLine[0], "X+"))
		byOffset, _ := strconv.Atoi(strings.TrimPrefix(buttonBLine[1], "Y+"))

		// Calculate button coordinates
		ax, ay := axOffset, ayOffset
		bx, by := bxOffset, byOffset

		// Store result
		results = append(results, Result{
			AX: ax, AY: ay,
			BX: bx, BY: by,
			PRIZE_X: prizeX, PRIZE_Y: prizeY,
		})
	}
	return results
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
