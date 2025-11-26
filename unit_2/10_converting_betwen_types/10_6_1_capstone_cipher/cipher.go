package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "I need help please"
	preparedCipherText := strings.ToUpper(strings.ReplaceAll(cipherText, " ", ""))
	cipherTextLen := len(preparedCipherText)
	keyWord := "GOLANG"
	keyWordLen := len(keyWord)
	message := ""
	keyIndex := 0

	for i := 0; i < cipherTextLen; i++ {
		// A=0, B=1, ... Z=25
		c := preparedCipherText[i] - 'A'
		k := keyWord[keyIndex] - 'A'

		// cipher letter + key letter
		c = (c+k+26)%26 + 'A'
		message += string(c)

		keyIndex++
		keyIndex %= keyWordLen
	}

	fmt.Println(message)
}
