package main

import "fmt"

// Product -
type Product struct {
	Name string
	Length int
	Width int
	Height int
}

// Box -
type Box struct {
	Length int
	Width int
	Height int
}

func getBestBox(
	availableBoxes []Box, 
	products []Products,
	) Box {
	
		//TODO: Choose the presice box
		fmt.Println("test")
		return Box{}
}