package assets

import (
	"embed"
)

var fs embed.FS

// Scissors returns the scissors image as bytes.
func Scissors() ([]byte, error) {
	return fs.ReadFile("raw/scissors.png")
}

// CHCross returns the CH-cross image as bytes.
func CHCross() ([]byte, error) {
	return fs.ReadFile("raw/ch-cross.png")
}
