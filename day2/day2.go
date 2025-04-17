package main

import (
	"bufio"
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

	report_data := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // internally, it advances token based on sperator

		arr := strings.Split(scanner.Text(), " ")
		nested_slice := make([]int, 0)
		for i := range arr {

			val, err := strconv.Atoi(arr[i])
			if err != nil {
				// ... handle error
				panic(err)
			} else {
				nested_slice = append(nested_slice, val)
			}

		}
		report_data = append(report_data, nested_slice)
	}
	output := 0
	for _, row := range report_data {

		increasing := row[0] < row[len(row)-1] || row[1] < row[len(row)-1]
		prev_val := -1
		incorrect := 0
		row_incorrect := 0
		// nums_reset := 0
		for i, _ := range row {

			if increasing {
				//condition is not met

				for index, row_val := range row {
					if i == index {
						//println(i)
					} else if prev_val == -1 {
						prev_val = row_val
					} else if !(row_val > prev_val) || !(row_val-prev_val < 4) {
						incorrect++
					} else {
						prev_val = row_val
					}
				}
				if incorrect > 0 {
					row_incorrect++
				}
				incorrect = 0
				prev_val = -1
			} else {
				for index, row_val := range row {
					if i == index {
						//println(i)
					} else if prev_val == -1 {
						prev_val = row_val
					} else if !(row_val < prev_val) || !(prev_val-row_val < 4) {
						incorrect++
					} else {
						prev_val = row_val
					}
				}
				if incorrect > 0 {
					row_incorrect++
				}
				incorrect = 0
				prev_val = -1

			}
		}
		if row_incorrect < len(row) {
			output++

		}

	}
	print(output)

}
