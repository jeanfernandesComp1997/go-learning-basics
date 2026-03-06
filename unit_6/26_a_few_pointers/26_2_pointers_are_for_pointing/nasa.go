package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	bolden = "Maj. Gen. Charles F. Bolden Jr."
	fmt.Println(*administrator)

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)

	fmt.Println(administrator == major)

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major)

	fmt.Println(*administrator)

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	fmt.Println(bolden)

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	fmt.Println(&charles == &bolden)

	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
	fmt.Println(timmy.superpower)
	fmt.Println((*timmy).superpower)
}
