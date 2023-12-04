package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

var gameRegex, _ = regexp.Compile(`^Game (?P<Game>\d+)`)
var playRegex, _ = regexp.Compile(`(?P<Play>((\d+ (green|red|blue),? ?){1,3})(;|$|\n))`)
var diceRegex, _ = regexp.Compile(`(?P<Count>\d+) (?P<Color>red|green|blue)`)

type Sol struct {
	Sum int
}

func (s *Sol) IsPossible(roll map[string]int) bool {
	for color, count := range roll {
		if count > limits[color] {
			return false
		}
	}
	return true
}

func (s *Sol) Read(str string) {
	gameMatch := gameRegex.FindStringSubmatch(str)[gameRegex.SubexpIndex("Game")]
	game, _ := strconv.Atoi(gameMatch)
	validGame := true
	for _, match := range playRegex.FindAllStringSubmatch(str, -1) {
		play := match[playRegex.SubexpIndex("Play")]
		var roll = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, diceMatch := range diceRegex.FindAllStringSubmatch(play, -1) {
			count := diceMatch[diceRegex.SubexpIndex("Count")]
			color := diceMatch[diceRegex.SubexpIndex("Color")]
			roll[color], _ = strconv.Atoi(count)
		}
		if !s.IsPossible(roll) {
			validGame = false
			break
		}
	}
	if validGame {
		s.Sum += game
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	s := &Sol{}
	for scanner.Scan() {
		line := scanner.Text()
		s.Read(line)
	}

	log.Printf("Sum: %v", s.Sum)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
