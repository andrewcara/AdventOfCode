package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// color_map := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }
	output := 0
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		draw_color := make([][]string, 0)

		line := strings.Split(fileScanner.Text(), ":")
		// game_number := strings.Split(line[0], " ")[1]
		// game_number_int, _ := strconv.Atoi(game_number)
		draws := strings.Split(line[1], ";")

		for _, colour := range draws {
			draw_color = append(draw_color, strings.Split(colour, ","))
		}

		max_map := make(map[string]int)

		for i := 0; i < len(draw_color); i++ {
			for j := 0; j < len(draw_color[i]); j++ {

				num := strings.Split(draw_color[i][j], " ")
				value, ok := max_map[num[2]]
				colour_value, _ := strconv.Atoi(num[1])

				if ok {
					if value < colour_value {
						max_map[num[2]] = colour_value
					}
				} else {
					max_map[num[2]] = colour_value
				}
			}
		}
		cube_max := 0

		for _, v := range max_map {
			if cube_max == 0 {
				cube_max = v
			} else {
				cube_max *= v
			}
		}

		output += cube_max
	}
	println(output)
}
