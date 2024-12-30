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

	var Reset = "\033[0m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDont := regexp.MustCompile(`don't\(\).*?do\(\)`) // Returns everything between don't() and do() in "reverse"

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		match := reDont.ReplaceAllString(line, "")
		matches := reDont.FindAllStringIndex(line, -1)
		foundMatches := reMul.FindAllStringSubmatch(match, -1)

		//fmt.Println(foundMatches)

		for _, foundMatch := range foundMatches {
			//fmt.Println(Red + foundMatch[1] + foundMatch[2] + Reset)
			first, _ := strconv.Atoi(foundMatch[1])
			second, _ := strconv.Atoi(foundMatch[2])
			sum += first * second
		}

		// Build the colored line
		coloredLine := ""
		lastIndex := 0
		for _, match := range matches {
			// Append the text before the match
			coloredLine += Green + line[lastIndex:match[0]]
			// Append the matched text in red
			coloredLine += Red + line[match[0]:match[1]] + Reset
			// Update the last index
			lastIndex = match[1]
		}
		// Append the remaining text after the last match
		coloredLine += Green + line[lastIndex:] + Reset
		fmt.Println(coloredLine)
	}
	fmt.Println(sum)
}
