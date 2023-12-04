package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	var pattern = `\d`
	var regex, err = regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Error compiling regex: %s", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllString(line, -1)
		first := matches[0]
		last := matches[len(matches)-1]
		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		log.Printf("> %v - %d", line, num)
		sum += num
	}

	log.Printf("Sum: %v", sum)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
