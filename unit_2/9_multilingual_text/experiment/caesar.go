package main

import "fmt"

func main() {
	caesar()
	international()
}

func caesar() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 13
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func international() {
	message := "Hola Estación Espacial Internacional" // ROT13 cypher
	fmt.Println(len(message))

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
