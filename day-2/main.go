package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		var digits []int
		line := scanner.Text()
		fields := strings.Fields(line) // []string
		isAscending := false
		isDescending := false
		safe := true
		for _, field := range fields {
			digit, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			digits = append(digits, digit)
		}
		for i := 0; i < len(digits)-1; i++ {
			first := digits[i]
			second := digits[i+1]
			result := first - second
			switch math.Signbit(float64(result)) {
			case true:
				isAscending = true
				result = int(math.Abs(float64(result)))
			case false:
				isDescending = true
			}
			if isAscending && isDescending == true {
				safe = false
				break
			}
			if result < 1 || result > 3 {
				isAscending = false
				isDescending = false
				safe = false
				break
			}
		}
		if safe {
			counter++
		}
	}
	fmt.Println(counter)
}
