package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
	EndOfRow    = "/0"
	MaxDigits   = 4
	DoKeyword   = "do()"
	DontKeyword = "don't()"
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
	output := 0
	scanner := bufio.NewScanner(file)
	line_token := Token{}
	enabled := true

	for scanner.Scan() {
		row := scanner.Text()

		for i := 0; i < len(row); {
			switch {
			case string(row[i]) == "d":
				should_reset := stateReset(row, i, enabled)
				if should_reset {
					enabled = !enabled
				}
				i++

			case (!enabled):
				i++
				line_token.resetToken()
			case isMul(row, i):
				line_token.hasMul = true
				i += 3
			case string(row[i]) == "(":
				if line_token.hasMul {
					line_token.leftParenth = true
					i++
				} else {
					line_token.resetToken()
					i++
				}
			case isDigit(row, i):
				if line_token.hasComma {
					value, loops := findNumber(row, i)
					if value > 0 {
						i += loops
						line_token.second_digit = value
					} else {
						i++
						line_token.resetToken()
					}
				} else if line_token.leftParenth {
					value, loops := findNumber(row, i)

					if value > 0 {
						i += loops
						line_token.first_digit = value
					} else {
						i++
						line_token.resetToken()
					}
				} else {
					i++
					line_token.resetToken()
				}
			case string(row[i]) == ",":
				if line_token.first_digit > 0 {
					line_token.hasComma = true
					i++
				} else {
					line_token.resetToken()
					i++
				}
			case string(row[i]) == ")":
				if line_token.second_digit > 0 {
					output += (line_token.first_digit * line_token.second_digit)
					line_token.resetToken()
					i++
				} else {
					line_token.resetToken()
					i++
				}
			case string(row[i]) == " ":
				i++
			default:
				i++
				line_token.resetToken()

			}
		}
	}
	print(output)
}

func (curr_token *Token) resetToken() {
	*curr_token = Token{}
}

func stateReset(curr_row string, index int, curr_state bool) bool {
	if curr_state {
		return peekDont(curr_row, index)
	} else {
		return peekDo(curr_row, index)
	}
}

func peekDo(curr_row string, index int) bool {

	output := string(curr_row[index])

	for i := 0; i < (len(DoKeyword) - 1); {
		output += peek(curr_row, index+i)
		i++
	}
	return output == DoKeyword

}

func peekDont(curr_row string, index int) bool {
	output := string(curr_row[index])

	for i := 0; i < (len(DontKeyword) - 1); {
		output += peek(curr_row, index+i)
		i++
	}
	return output == DontKeyword
}

func isDigit(curr_row string, index int) bool {
	if _, err := strconv.Atoi(string(curr_row[index])); err == nil {
		return true
	}
	return false
}

func findNumber(curr_row string, index int) (int, int) {
	str_digit := string(curr_row[index])
	loops := 0

	for {
		val := peek(curr_row, index+loops)
		if _, err := strconv.Atoi(val); err == nil {
			str_digit += val
			loops++
		} else {
			break
		}
	}
	if len(str_digit) < MaxDigits {
		value, _ := strconv.Atoi(str_digit)
		return value, loops + 1
	}
	return 0, 0

}

func isMul(curr_row string, index int) bool {
	if string(curr_row[index]) == "m" && peek(curr_row, index) == "u" && peek(curr_row, index+1) == "l" {
		return true
	}
	return false
}

func peek(curr_row string, index int) string {
	if index+1 >= len(curr_row)-1 {
		return EndOfRow
	}
	return string(curr_row[index+1])
}
