package main

import (
	"fmt"
	"unique-codes-generator/generator"
)

func main() {
	const number = 100
	codes, _ := generator.GenIdsWithSonyFlake(100)
	fmt.Printf("%v", len(codes))
	fmt.Printf("%v", codes)
}
