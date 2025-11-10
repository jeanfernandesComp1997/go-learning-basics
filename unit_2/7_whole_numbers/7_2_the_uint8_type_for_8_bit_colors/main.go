package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0, 141, 213
	fmt.Println(red, green, blue)

	var redHex, greenHex, blueHex uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x\n", redHex, greenHex, blueHex)
	fmt.Printf("color: #%02x%02x%02x;", redHex, greenHex, blueHex)
}
