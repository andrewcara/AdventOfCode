package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	X_POS, Y_POS, X_VEL, Y_VEL int
}
type Min struct {
	max  int
	time int
}
type Key struct {
	X, Y int
}

func main() {
	input := getInput()
	robots := getRobots(input)
	width := 101
	height := 103
	min_obj := Min{max: 0, time: 1}

	for i := 1; i < 150000; i++ {
		q1, q2, q3, q4 := 0, 0, 0, 0
		robot_map := make(map[Key]int)
		for _, val := range robots {
			x, y := getPositionAfterTime(val, i, width, height)

			if (x < width/2) && (y < height/2) {
				q1++
			} else if (x > width/2) && (y < height/2) {
				q2++
			} else if (x < width/2) && (y > height/2) {
				q3++
			} else if (x > width/2) && (y > height/2) {
				q4++
			}
			robot_map[Key{x, y}] = 1
		}
		if len(robot_map) > min_obj.max {
			min_obj.max = len(robot_map)
			min_obj.time = i
			fmt.Printf("----- POTENTIAL PICTURE @ %d seconds -----", i)

		}
	}
	println(min_obj.time)

}

func getPositionAfterTime(robot Robot, time, x_dimension, y_dimension int) (int, int) {
	// Calculate positions after time
	x_after_time := (time * robot.X_VEL) + robot.X_POS
	y_after_time := (time * robot.Y_VEL) + robot.Y_POS

	// Handle wrapping using modular arithmetic
	x := ((x_after_time % x_dimension) + x_dimension) % x_dimension
	y := ((y_after_time % y_dimension) + y_dimension) % y_dimension
	return x, y
}

func getRobots(input string) []Robot {
	robots := make([]Robot, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, row := range lines {
		position := strings.Split(strings.Split(row, " ")[0], ",")
		velocity := strings.Split(strings.Split(row, " ")[1], ",")

		x_pos, _ := strconv.Atoi(strings.Split(position[0], "=")[1])
		y_pos, _ := strconv.Atoi(position[1])

		x_vel, _ := strconv.Atoi(strings.Split(velocity[0], "=")[1])
		y_vel, _ := strconv.Atoi(velocity[1])

		robots = append(robots, Robot{X_POS: x_pos, Y_POS: y_pos, X_VEL: x_vel, Y_VEL: y_vel})

	}
	return robots
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
