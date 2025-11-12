package main

import "fmt"

func main() {
	c := 'a'
	c += 3
	fmt.Printf("%c\n", c)

	d := 'g'
	d = d - 'a' + 'A'
	fmt.Printf("%c\n", d)

	rot13()
}

func rot13() {
	message := "uv vagreangvbany fcnpr fgngvba" // ROT13 cypher
	fmt.Println(len(message))

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
