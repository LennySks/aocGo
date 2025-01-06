package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	word := "XMAS"
	reverseWord := Reverse(word)
	wordLength := len(word)
	count := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// horizontal
	for _, line := range lines {
		for x := 0; x <= len(line)-wordLength; x++ {
			if line[x:x+wordLength] == word || line[x:x+wordLength] == reverseWord {
				count++
			}
		}
	}

	// vertical

	for col := 0; col < len(lines[0]); col++ {
		for row := 0; row <= len(lines)-wordLength; row++ {
			verticalWord := ""
			for k := 0; k < wordLength; k++ {
				verticalWord += string(lines[row+k][col])
			}
			if verticalWord == word || verticalWord == reverseWord {
				count++
			}
		}
	}

	// diagonal top left to bottom right

	for row := 0; row <= len(lines)-wordLength; row++ {
		for col := 0; col <= len(lines[0])-wordLength; col++ {
			diagonalWord := ""
			for k := 0; k < wordLength; k++ {
				diagonalWord += string(lines[row+k][col+k])
			}
			if diagonalWord == word || diagonalWord == reverseWord {
				count++
			}
		}
	}

	// diagonal top right to bottom left

	for row := 0; row <= len(lines)-wordLength; row++ {
		for col := wordLength - 1; col < len(lines[0]); col++ {
			diagonalWord := ""
			for k := 0; k < wordLength; k++ {
				diagonalWord += string(lines[row+k][col-k])
			}
			if diagonalWord == word || diagonalWord == reverseWord {
				count++
			}
		}
	}

	fmt.Println(count)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
