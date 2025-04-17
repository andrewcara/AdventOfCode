package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
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
	output := 0
	for scanner.Scan() {
		row := scanner.Text()
		list, total := format_row(row)
		current_total := list[0]
		start_index := 1
		if perform_operation(list, total, current_total, start_index) {
			output += total
		}

	}
	println(output)
}

func perform_operation(values []int, total int, current_total int, index int) bool {
	//println(current_total, index, values[index])
	if current_total == total && (len(values) == index) {
		return true
	}
	if len(values) == index {
		return false
	}
	if perform_operation(values, total, current_total+values[index], index+1) {
		return true
	}
	if perform_operation(values, total, current_total*values[index], index+1) {
		return true
	}
	if perform_operation(values, total, concatenate_nums(current_total, values[index]), index+1) {
		return true
	}

	return false
}

func format_row(input_string string) ([]int, int) {

	total, _ := strconv.Atoi(strings.Split(input_string, ":")[0])
	trimmed_str := strings.TrimLeft((strings.Split(input_string, ":")[1]), " ")
	row_str := strings.Split(trimmed_str, " ")
	row := make([]int, 0)
	for _, val := range row_str {
		cast_val, _ := strconv.Atoi(val)
		row = append(row, cast_val)
	}
	return row, total
}

func concatenate_nums(number1 int, number2 int) int {
	var b bytes.Buffer
	str1 := strconv.Itoa(number1)
	str2 := strconv.Itoa(number2)
	b.WriteString(str1)
	b.WriteString(str2)
	concatenated_str := b.String()

	concatenated_num, _ := strconv.Atoi(concatenated_str)

	return concatenated_num
}
