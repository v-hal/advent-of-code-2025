package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	var min_position = 0
	var max_position = 99
	var full_rotation = max_position - min_position + 1
	var rotations = read_rotations("input.txt")
	var current_position = 50
	var zero_position_visited_count = 0

	for _, rotation := range rotations {
		current_position = (current_position + rotation) % full_rotation
		if current_position == 0 {
			zero_position_visited_count++
		}
	}
	fmt.Println(zero_position_visited_count)
}
