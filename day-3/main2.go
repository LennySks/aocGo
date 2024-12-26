package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	//"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		foundMatches := r.FindAllStringSubmatch(line, -1)
		for _, match := range foundMatches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}
	fmt.Printf("Total sum: %d", sum)
}
