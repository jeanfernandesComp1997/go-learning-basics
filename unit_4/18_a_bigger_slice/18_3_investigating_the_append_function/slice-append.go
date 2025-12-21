package main

import "fmt"

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dwarfs3[1] = "Pluto!"
	dwarfs1[0] = "Ceres!"

	fmt.Printf("%v\n%v\n%v\n", dwarfs1, dwarfs2, dwarfs3)
}
