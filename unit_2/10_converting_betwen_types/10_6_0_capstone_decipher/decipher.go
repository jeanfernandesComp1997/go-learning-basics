package main

import "fmt"

func main() {
	bookSolution()
}

func bookSolution() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	cipherTextLen := len(cipherText)
	keyWord := "GOLANG"
	keyWordLen := len(keyWord)
	message := ""
	keyIndex := 0

	for i := 0; i < cipherTextLen; i++ {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A'
		k := keyWord[keyIndex] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A'
		message += string(c)

		keyIndex++
		keyIndex %= keyWordLen
	}

	fmt.Println(message)
}
