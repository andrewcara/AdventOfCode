package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	counter := 0
	val1_list := make([]int, 0)
	val2_list := make([]int, 0)

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

	for scanner.Scan() { // internally, it advances token based on sperator
		arr := strings.Split(scanner.Text(), " ")

		cast_val1, err := strconv.Atoi(arr[0])
		if err != nil {
			// ... handle error
			panic(err)
		}

		cast_val2, err := strconv.Atoi(arr[len(arr)-1])
		if err != nil {
			// ... handle error
			panic(err)
		}
		val1_list = append(val1_list, cast_val1)
		val2_list = append(val2_list, cast_val2)

	}

	// sort.Slice(val1_list, func(i, j int) bool {
	// 	return val1_list[i] < val1_list[j]
	// })

	// sort.Slice(val2_list, func(i, j int) bool {
	// 	return val2_list[i] < val2_list[j]
	// })
	dict := map[int]int{}

	for i := 0; i < len(val1_list); i++ {

		dict[val2_list[i]]++
	}

	for j := 0; j < len(val1_list); j++ {
		counter += (dict[val1_list[j]] * val1_list[j])
	}

	print(counter)
}
