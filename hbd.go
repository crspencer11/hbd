package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	width       = 70  // Terminal width
	height      = 20  // Terminal height
	snowflakes  = "*" // Symbols for snowflakes
	clearScreen = "\033[H\033[J" // ANSI escape code to clear the terminal
)

// Snowflake structure to track position
type Snowflake struct {
	x, y int
	char rune
}

// Function to initialize snowflakes
func initSnowflakes(count int) []Snowflake {
	snow := make([]Snowflake, count)
	for i := range snow {
		snow[i] = Snowflake{
			x:    rand.Intn(width),
			y:    rand.Intn(height),
			char: rune(snowflakes[rand.Intn(len(snowflakes))]),
		}
	}
	return snow
}

// Function to update snowflake positions
func updateSnowflakes(snow []Snowflake) {
	for i := range snow {
		snow[i].y++ // Move snowflake down
		if snow[i].y >= height { // If snowflake reaches the bottom, reset to top
			snow[i].y = 0
			snow[i].x = rand.Intn(width)
			snow[i].char = rune(snowflakes[rand.Intn(len(snowflakes))])
		}
	}
}

// Function to render the snowflakes and moving message
func renderFrame(snow []Snowflake, message string, offset int) {
	fmt.Print(clearScreen) // Clear the screen

	// Create a 2D grid for rendering
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = ' ' // Default to empty space
		}
	}

	// Place snowflakes into the grid
	for _, flake := range snow {
		if flake.y < height && flake.x < width {
			grid[flake.y][flake.x] = flake.char
		}
	}

	// Place the horizontally moving birthday message
	startX := offset % (width + len(message))
	for i, char := range message {
		pos := (startX + i) % width // Wrap the message horizontally
		if pos >= 0 && pos < width {
			grid[height/2][pos] = char
		}
	}

	// Print the grid
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get name from arguments or default
	name := "Friend"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message := fmt.Sprintf(" 🎉 Happy Birthday %s! 🎉 ", name) // Message with padding

	snow := initSnowflakes(40)

	offset := 0
	for i := 0; i < 120; i++ { // 300 animation frames
		renderFrame(snow, message, offset)
		updateSnowflakes(snow)
		offset++ // Move the message horizontally
		time.Sleep(90 * time.Millisecond)
	}

	fmt.Println("\n")
	fmt.Println("🎊 Enjoy your special day! 🎊")
}

