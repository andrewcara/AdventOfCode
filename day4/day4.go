package main

import (
	"bufio"
	"log"
	"os"
)

type Token struct {
	hasMul       bool
	leftParenth  bool
	first_digit  int
	hasComma     bool
	second_digit int
	rightParenth bool
}

const (
	Keyword = "XMAS"
)

func main() {
	file, err := os.Open("input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	row_len, row_height := get_dimensions(scanner)

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	scanner = bufio.NewScanner(file)

	// Adjusting dimensions for `buffered_slice`
	padding := (len(Keyword)) * 2
	buffered_slice := make([][]string, row_height+padding)

	for row := range buffered_slice {
		buffered_slice[row] = make([]string, row_len+padding+1)

		for col := range buffered_slice[row] {
			if row > len(Keyword)-1 && row < len(buffered_slice)-len(Keyword) {
				if col < len(Keyword) || col > len(buffered_slice)-len(Keyword) && col < len(buffered_slice)+1 {
					buffered_slice[row][col] = "."
				}

			} else {
				buffered_slice[row][col] = "."
			}
		}

	}

	height := 0
	for scanner.Scan() {
		row := scanner.Text()
		for i := 0; i < len(row); i++ { // Increment i in the loop
			buffered_slice[height+len(Keyword)][i+len(Keyword)] = string(row[i])
		}
		height++ // Move to the next row
	}
	output := 0
	for row_index, row := range buffered_slice {

		for col_index, val := range row {

			if val == "A" {
				if check_xmas(buffered_slice, row_index, col_index) {
					output++
				}
			}
		}
	}
	print(output)
}

func get_dimensions(scanner *bufio.Scanner) (int, int) {
	row_len := 0
	row_height := 0
	for scanner.Scan() {

		row := scanner.Text()
		row_len = len(row)
		row_height++

	}
	return row_len, row_height
}

func check_direction(array [][]string, row int, col int) int {

	//check left and right
	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}
	output := 0
	wordLen := len(Keyword)
	for _, dir := range directions {
		dRow, dCol := dir[0], dir[1]
		matched := true

		for i := 0; i < wordLen; i++ {
			r := row + dRow*i
			c := col + dCol*i

			// Check boundaries
			if r < 0 || r >= len(array) || c < 0 || c >= len(array[0]) {
				matched = false
				break
			}

			// Check character match
			if array[r][c] != string(Keyword[i]) {
				matched = false
				break
			}
		}

		if matched {
			output++
		}
	}
	return output
}
func check_xmas(array [][]string, row int, col int) bool {

	if ((string(array[row+1][col+1]) == "S" && string(array[row-1][col-1]) == "M") || (string(array[row+1][col+1]) == "M" && string(array[row-1][col-1]) == "S")) && (string(array[row-1][col+1]) == "S" && string(array[row+1][col-1]) == "M" || string(array[row-1][col+1]) == "M" && string(array[row+1][col-1]) == "S") {

		return true
	}

	return false
}
