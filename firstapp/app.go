package main

import (
	"fmt"

	"rsc.io/quote"

	"encoding/json"
)

func main() {
	fmt.Println(quote.Go())

	a := json.Number(2)

	fmt.Println(a)

	message := greetings.Hello("Mumy")
	fmt.Println(message)
}
