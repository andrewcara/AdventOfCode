package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	Y, X int
}

type Character struct {
	Position Position
	Char     string
}

func (c Character) addDirection(Current_Direction Position, mult int) Position {

	return Position{Y: c.Position.Y + (Current_Direction.Y * mult), X: c.Position.X + (Current_Direction.X * mult)}
}

type Robot struct {
	Position          Position
	Current_Direction Position
}

var brackets = map[string]string{
	"[": "]",
	"]": "[",
}

func (r Robot) addDirection(mult int) Position {

	return Position{Y: r.Position.Y + (r.Current_Direction.Y * mult), X: r.Position.X + (r.Current_Direction.X * mult)}
}

func main() {
	input := getInput()

	position_map, robot := getMapInput(input)
	move_sequence := getMoveInput(input)
	//fmt.Println(position_map, robot)
	output := part2(input, position_map, &robot, move_sequence)

	fmt.Println(output, robot.Position)

}

func part2(input string, position_map map[Position]string, robot *Robot, move_sequence string) int {
	output := 0
	moves := map[string]Position{"^": {-1, 0}, "<": {0, -1}, ">": {0, 1}, "v": {1, 0}}
	fmt.Println(robot.Position)
	for _, char := range move_sequence {
		robot.Current_Direction = moves[string(char)]

		if position_map[robot.addDirection(1)] == "." {
			position_map[robot.Position] = "."
			robot.Position = robot.addDirection(1)

		} else if position_map[robot.addDirection(1)] == "[" || position_map[robot.addDirection(1)] == "]" {

			if string(char) == "<" || string(char) == ">" {
				check_side(position_map, robot)
			}

			if string(char) == "^" || string(char) == "v" {
				check_vert(position_map, robot)
			}

		}
	}
	for k, v := range position_map {

		if k.Y == 3 {
			fmt.Println(k, v)
		}

		if v == "O" {
			output += (k.Y * 100) + k.X
		}
	}

	return 0
}

func check_vert(position_map map[Position]string, robot *Robot) {
	current_tier := make([]Character, 0)
	total_brackets := make([]Character, 0)
	bracket := position_map[robot.addDirection(1)]
	current_tier = append(current_tier, Character{Position: robot.addDirection(1), Char: bracket})
	current_tier = append(current_tier, Character{Position: closingBracketPos(bracket, robot.addDirection(1)), Char: brackets[bracket]})
	total_brackets = append(total_brackets, current_tier...)
	end_index := 2
	start_index := 0
	tier := 2
	for len(current_tier) > 0 {
		current_tier := make([]Character, 0)
		for _, char := range total_brackets[start_index:end_index] {

			if position_map[char.addDirection(robot.Current_Direction, tier)] == "[" {
				current_tier = append(current_tier, Character{Position: char.addDirection(robot.Current_Direction, tier), Char: "["})
				current_tier = append(current_tier, Character{Position: closingBracketPos(bracket, char.addDirection(robot.Current_Direction, tier)), Char: "]"})
				end_index += 2
			} else if position_map[char.addDirection(robot.Current_Direction, tier)] == "[" {
				current_tier = append(current_tier, Character{Position: char.addDirection(robot.Current_Direction, tier), Char: "]"})
				current_tier = append(current_tier, Character{Position: closingBracketPos(bracket, char.addDirection(robot.Current_Direction, tier)), Char: "["})
			}

		}
	}
}

func closingBracketPos(bracket string, bracket_pos Position) Position {
	if bracket == "]" {
		return Position{Y: bracket_pos.Y, X: bracket_pos.X - 1}
	}
	return Position{Y: bracket_pos.Y, X: bracket_pos.X + 1}
}

func check_side(position_map map[Position]string, robot *Robot) {
	i := 3
	start_bracket := position_map[robot.addDirection(1)]
	//println("here")

	for position_map[robot.addDirection(i)] != "." && position_map[robot.addDirection(i)] != "#" {
		i += 2

	}
	terminal_char := position_map[robot.addDirection(i)]

	//fmt.Println(position_map[robot.addDirection(i)], i, start_bracket, brackets[start_bracket])

	if terminal_char == "." {
		position_map[robot.addDirection(1)] = "."

		for i > 1 {
			fmt.Println(position_map[robot.addDirection(i-2)], i, brackets[start_bracket], start_bracket)

			position_map[robot.addDirection(i)] = brackets[start_bracket]
			position_map[robot.addDirection(i-1)] = start_bracket
			i -= 2
		}
		robot.Position = robot.addDirection(1)

	}
}

func part1(input string, position_map map[Position]string, robot Robot, move_sequence string) int {
	output := 0
	moves := map[string]Position{"^": {-1, 0}, "<": {0, -1}, ">": {0, 1}, "v": {1, 0}}
	for _, char := range move_sequence {
		robot.Current_Direction = moves[string(char)]

		if position_map[robot.addDirection(1)] == "." {
			position_map[robot.Position] = "."
			robot.Position = robot.addDirection(1)

		} else if position_map[robot.addDirection(1)] == "O" {
			i := 2
			for position_map[robot.addDirection(i)] != "." && position_map[robot.addDirection(i)] != "#" {
				i++
			}
			terminal_char := position_map[robot.addDirection(i)]

			if terminal_char == "." {
				position_map[robot.addDirection(1)] = "."
				position_map[robot.addDirection(i)] = "O"
				robot.Position = robot.addDirection(1)
			}
		}
	}
	for k, v := range position_map {
		fmt.Println(k, v)

		if v == "O" {
			output += (k.Y * 100) + k.X
		}
	}
	return output
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

func getMapInput(input string) (map[Position]string, Robot) {

	input = strings.Split(input, "\n\n")[0]
	input = strings.ReplaceAll(input, "\r\n", "\n")
	position_map := make(map[Position]string)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	robot := Robot{Position: Position{0, 0}, Current_Direction: Position{0, 0}}
	//fmt.Println(lines)
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			switch {
			case string(line[j]) == "@":
				robot.Position = Position{Y: i, X: j * 2}
				position_map[Position{Y: i, X: j * 2}] = "."
				position_map[Position{Y: i, X: j*2 + 1}] = "."
			case string(line[j]) == "O":
				position_map[Position{Y: i, X: j * 2}] = "["
				position_map[Position{Y: i, X: j*2 + 1}] = "]"
			default:
				position_map[Position{Y: i, X: j * 2}] = string(line[j])
				position_map[Position{Y: i, X: j*2 + 1}] = string(line[j])
			}
		}
	}
	return position_map, robot
}

func getMoveInput(input string) string {
	prs_output := strings.Split(input, "\n\n")
	lines := strings.Split(prs_output[1], "\n")
	var buffer bytes.Buffer
	for _, line := range lines {
		buffer.WriteString((line))
	}
	return buffer.String()

}
