package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	person := Person{
		Name: "Ryan",
		Age:  20,
	}

	fmt.Printf(" %+v \n", person)

	nextPerson := NewPerson("Ayo", 26)

	fmt.Printf(" %+v \n", nextPerson)

	person.changeName("Ryan Van")

	fmt.Printf(" %+v \n", person)

	a := 7
	b := &a
	*b = 9

	fmt.Printf(" %+v \n", b)
	fmt.Printf(" %+v \n", *b)
	fmt.Printf(" %+v \n", a)

	mySlice := []int{1, 2, 3}

	for idx := range mySlice {
		mySlice[idx]++
	}

	fmt.Println(mySlice)
}
