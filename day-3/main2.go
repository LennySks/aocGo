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
	file, err := os.Open("./example2.txt")
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

		coloredLine := ""
		lastIndex := 0
		for _, match := range matches {
			coloredLine += Green + line[lastIndex:match[0]]
			coloredLine += Red + line[match[0]:match[1]] + Reset
			lastIndex = match[1]
		}
		coloredLine += Green + line[lastIndex:] + Reset
		fmt.Println(coloredLine)

		for _, foundMatch := range foundMatches {
			first, _ := strconv.Atoi(foundMatch[1])
			second, _ := strconv.Atoi(foundMatch[2])
			sum += first * second
		}
	}
	fmt.Println(sum)
}
