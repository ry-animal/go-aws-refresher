package main

import (
	"fmt"
	"go-aws/imports"
)

func main() {
	newTicket := imports.Ticket{
		ID:    1,
		Event: "go-aws",
	}
	fmt.Printf(" %+v \n", newTicket)

	newTicket.PrintEvent()
}
