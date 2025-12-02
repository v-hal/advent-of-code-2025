package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const min_position = 0
const max_position = 99
const start_position = 50
const full_rotation = max_position - min_position + 1

func read_rotations(fileName string) []int {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err, "reading the input file failed")
	}
	var rotations = []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		var direction = string(instruction[0])
		var clicks, err = strconv.Atoi(strings.TrimPrefix(instruction, direction))
		if err != nil {
			log.Fatal(err, "parsing the clicks failed")
		}
		if direction == "R" {
			rotations = append(rotations, clicks)
		} else if direction == "L" {
			rotations = append(rotations, -clicks)
		} else {
			log.Fatal("Unknown direction")
		}
	}
	return rotations
}

func mod(a, b int) int {
	r := a % b
	if r < 0 {
		r += b
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(rotations []int) int {
	var current_position = start_position
	var stopped_at_zero_count = 0

	for _, rotation := range rotations {
		current_position = mod(current_position+rotation, full_rotation)
		if current_position == 0 {
			stopped_at_zero_count++
		}
	}
	return stopped_at_zero_count
}

func part2(rotations []int) int {
	var current_position = start_position
	var zero_position_visited_count = 0

	for _, rotation := range rotations {
		var full_rotations = rotation / full_rotation
		var rotation_without_full_rotations = rotation - (full_rotations * full_rotation)
		var distance_to_zero_left = current_position + rotation_without_full_rotations
		var times_zero_visited = abs(full_rotations)

		if current_position != 0 && distance_to_zero_left <= min_position || distance_to_zero_left > max_position {
			times_zero_visited++
		}
		current_position = mod(current_position+rotation, full_rotation)
		zero_position_visited_count += times_zero_visited

	}
	return zero_position_visited_count
}

func main() {
	var rotations = read_rotations("input.txt")
	fmt.Println(part1(rotations))
	fmt.Println(part2(rotations))
}
