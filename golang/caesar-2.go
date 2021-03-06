package main

import (
	"flag"
	"fmt"
	"unicode"
)

func encode(char rune, distance int) rune {
	if unicode.IsLetter(char) {
		base := 'a'

		if distance < 0 {
			base = 'z'
		}

		if unicode.IsUpper(char) {
			base = unicode.ToUpper(base)
		}

		return ((char - base + rune(distance)) % 26) + base
	}

	return char
}

func encodeText(text string, distance int) string {
	cipher := make([]rune, len(text))

	for index, char := range text {
		cipher[index] = encode(char, distance)
	}

	return string(cipher)
}

func main() {
	decodeFlag := flag.Bool("d", false, "decode")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
	}

	text := flag.Arg(0)

	if *decodeFlag {
		fmt.Println(encodeText(text, -3))
	} else {
		fmt.Println(encodeText(text, 3))
	}
}
