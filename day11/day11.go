package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Assume getInput and getStonesMap are defined elsewhere.
	input := getInput()
	stones, length := getStonesMap(input) // `stones` is a map[string]int, and `length` is an int.

	for i := 0; i < 75; i++ {
		mem := make(map[string]int)

		// Iterate over the stones map.
		for key, value := range stones {
			if value != 0 {
				newValues := getStone(key)
				stones[key] = 0
				for _, newVal := range newValues {
					mem[newVal] += value
				}
				if len(newValues) > 1 {
					length += value
				}
			}
		}
		for key, value := range mem {
			stones[key] = value
		}
	}

	// Output the final state of the `stones` map.
	fmt.Println("Final Stones:", stones)
	fmt.Println("Total Length:", length)
}

func getStone(input string) []string {
	output := make([]string, 0)

	if input == "0" {
		output = append(output, "1")
	} else if len(input)%2 == 0 {
		first_half := input[0 : (len(input))/2]
		second_half := input[(len(input))/2:]
		second_half = stripZeros(second_half)
		output = append(output, first_half)
		output = append(output, second_half)
	} else {
		cast_val, _ := strconv.Atoi(string(input))
		cast_val *= 2024
		t := strconv.Itoa(cast_val)
		output = append(output, t)
	}
	return output

}

func stripZeros(input string) string {
	for i := range input {
		if string(input[i]) != "0" {
			return input[i:]
		}
	}
	return "0"
}

func getStonesMap(input string) (map[string]int, int) {
	stones := strings.Split(input, " ")
	mem := make(map[string]int)
	for _, val := range stones {
		mem[val] += 1
	}
	return mem, len(mem)
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
