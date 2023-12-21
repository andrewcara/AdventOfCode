package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parser(string_line string) (fistnumber int, secondnumber int) {
	numbers := map[string]int{

		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	i := 0
	num_in_str := make([]int, 0)
	current_str := ""
	for i < len(string_line) {

		j := i
		for j < len(string_line) {
			current_str += string(string_line[j])
			value, ok := numbers[current_str]

			if result, err := strconv.Atoi(current_str); err == nil {
				num_in_str = append(num_in_str, result)
			} else if ok {
				num_in_str = append(num_in_str, value)
			}
			j++
		}
		current_str = ""
		i++
	}
	return num_in_str[0], num_in_str[len(num_in_str)-1]
}

func main() {

	output := 0

	// first_number, second_number := parser("8qxbbcxjvppeightwot")
	// output += ((first_number * 10) + second_number)
	// println(output)
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		first_number, second_number := parser(string(fileScanner.Text()))
		output += ((first_number * 10) + second_number)
		println(output)
	}

}
