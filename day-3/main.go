package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	//"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	r2, _ := regexp.Compile("[0-9]+,[0-9]+")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		foundMatches := r.FindAllString(line, -1)
		for i := 0; i < len(foundMatches); i++ {
			foundDigits := r2.FindAllString(foundMatches[i], -1)
			for _, j := range foundDigits {
				fmt.Printf("Found digits pair: %s\n", j)
				newLine := strings.Split(j, ",")
				num1, err1 := strconv.Atoi(newLine[0])
				num2, err2 := strconv.Atoi(newLine[1])
				if err1 != nil || err2 != nil {
					log.Fatalf("Error converting to integer: %v, %v", err1, err2)
				}
				sum += num1 * num2
			}
		}
	}
	fmt.Printf("Total sum: %d", sum)
}
