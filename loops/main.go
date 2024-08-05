package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := []string{}
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")

	animals = slices.Delete(animals, 1, 2)

	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	for index, value := range animals {
		fmt.Printf("animals[%d] = %s\n", index, value)
	}

	for value := range 10 {
		fmt.Println(value)
	}

	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
}
