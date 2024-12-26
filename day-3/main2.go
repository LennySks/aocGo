package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		digits := parseDigits(fields)

		if isSafe(digits) || canBeMadeSafe(digits) {
			counter++
		}
	}
	fmt.Println(counter)
}

func isSafe(digits []int) bool {
	isAscending, isDescending := false, false
	for i := 0; i < len(digits)-1; i++ {
		result := digits[i] - digits[i+1]

		if result > 0 {
			isDescending = true
		} else if result < 0 {
			isAscending = true
			result = -result
		}

		if result < 1 || result > 3 || (isAscending && isDescending) {
			return false
		}
	}
	return true
}

func parseDigits(fields []string) []int {
	var digits []int
	for _, field := range fields {
		digit, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, digit)
	}
	return digits
}

func canBeMadeSafe(digits []int) bool {
	for i := 0; i < len(digits); i++ {
		modified := append([]int{}, digits[:i]...)
		modified = append(modified, digits[i+1:]...)
		if isSafe(modified) {
			return true
		}
	}
	return false
}
