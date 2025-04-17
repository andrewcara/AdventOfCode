package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	input := getInput()
	disk, empty, files := createDisk(input)

	for i := len(files) - 1; i > -1; i-- {
		for j, memory := range empty {

			if files[i].length <= memory.length && files[i].start_index > memory.start_index {
				taken_mem := 0
				for j := 0; j < files[i].length; j++ {
					disk[memory.start_index+j] = files[i].value
					disk[files[i].start_index+j] = -1
					taken_mem++
				}
				empty[j].length -= taken_mem
				empty[j].start_index += taken_mem
				break
			}
		}
	}
	output := 0

	for i, val := range disk {
		if val == -1 {
			continue
		} else {
			output += (i * val)
		}

	}

	fmt.Println(output)
}

type EmptyMemory struct {
	start_index int
	length      int
}
type Files struct {
	start_index int
	length      int
	value       int
}

func part1() {
	input := getInput()
	disk, _, _ := createDisk(input)

	left_pointer := 0
	right_pointer := len(disk) - 1

	for left_pointer < right_pointer {
		// Skip non-empty positions from left
		for left_pointer < right_pointer && disk[left_pointer] != -1 {
			left_pointer++
		}
		// Skip empty positions from right
		for left_pointer < right_pointer && disk[right_pointer] == -1 {
			right_pointer--
		}
		// Swap if conditions are met
		if left_pointer < right_pointer {
			disk[left_pointer] = disk[right_pointer]
			disk[right_pointer] = -1
			left_pointer++
			right_pointer--
		}
	}
	output := 0

	for i, val := range disk {
		if val == -1 {
			break
		}
		output += (i * val)
	}
	println(output)
}

func createDisk(input string) ([]int, []EmptyMemory, []Files) {
	disk := make([]int, 0)
	counter := 0
	empty_slots := make([]EmptyMemory, 0)
	files := make([]Files, 0)

	for index := range input {
		cast_val, _ := strconv.Atoi(string(input[index]))
		start_index := len(disk)
		if index%2 == 0 {

			for i := 0; i < cast_val; i++ {
				disk = append(disk, counter)

			}
			files = append(files, Files{start_index: start_index, length: cast_val, value: counter})
			counter++
		} else {
			for i := 0; i < cast_val; i++ {
				disk = append(disk, -1)
			}
			if cast_val > 0 {
				empty_slots = append(empty_slots, EmptyMemory{start_index: start_index, length: cast_val})
			}

		}
	}
	return disk, empty_slots, files
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
