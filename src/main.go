package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/inancgumus/screen"
)

type Input struct {
	text      string
	shift     int
	direction string
}

var alphabet [56]rune
var user Input

func art() {
	logo := figure.NewColorFigure("Caesar Cipher", "", "blue", true)
	logo.Print()
}

func cipher(text string, shift int, direction string) string {
	final := ""

	if direction == "decode" {
		shift *= -1
	}

	for _, char := range text {

		for pos, letter := range alphabet {

			if char == letter {
				final += string(alphabet[shift+pos])
			}
		}
	}

	return final
}

func populateAlphabet() {
	alphabet[0] = 'a'
	alphabet[1] = 'b'
	alphabet[2] = 'c'
	alphabet[3] = 'd'
	alphabet[4] = 'e'
	alphabet[5] = 'f'
	alphabet[6] = 'g'
	alphabet[7] = 'h'
	alphabet[8] = 'i'
	alphabet[9] = 'j'
	alphabet[10] = 'k'
	alphabet[11] = 'l'
	alphabet[12] = 'm'
	alphabet[13] = 'n'
	alphabet[14] = 'o'
	alphabet[15] = 'p'
	alphabet[16] = 'q'
	alphabet[17] = 'r'
	alphabet[18] = 's'
	alphabet[19] = 't'
	alphabet[20] = 'u'
	alphabet[21] = 'v'
	alphabet[22] = 'w'
	alphabet[23] = 'x'
	alphabet[24] = 'y'
	alphabet[25] = 'z'
}

func run() {

	populateAlphabet()

	should_end := true

	var text string
	var shift int
	var direction string
	var restart string

	for should_end {

		screen.Clear()
		screen.MoveTopLeft()
		art()
		fmt.Print("Type 'encode' to encrypt, type 'decode' to decrypt: ")
		fmt.Scanln(&direction)
		fmt.Print("Type your message: ")
		fmt.Scanln(&text)
		fmt.Print("Type the shift number: ")
		fmt.Scanln(&shift)

		shift = shift % 26

		user = Input{
			text:      text,
			shift:     shift,
			direction: direction,
		}

		fmt.Printf("Here's the " + user.direction + " result: " + cipher(user.text, user.shift, user.direction))

		fmt.Println("")
		fmt.Println("Type 'yes' if you want to go again. Otherwise type 'no'. ")
		fmt.Scanln(&restart)

		if restart == "no" {
			fmt.Println("")
			fmt.Println("Goodbye! S2")
			break
		}

	}

}

func main() {
	run()
}
