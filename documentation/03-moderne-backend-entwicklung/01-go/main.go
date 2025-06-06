package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/thenativeweb/techlounge/auth"
	"github.com/thenativeweb/techlounge/calculator"
)

func main() {
	fmt.Println(uuid.New())

	a := float64(23)
	b := float64(42)

	c := calculator.Add(a, b)
	fmt.Println(c)

	result, err := calculator.Divide(a, b)
	if err != nil {
		// panic(err)

		if err == calculator.ErrDivisionByZero {
			fmt.Println("Division by zero error has happened.")
		} else {
			fmt.Println("An unknown error has happened.")
		}

		os.Exit(1)
	}

	fmt.Println(result)

	user := auth.NewUser("Jane Doe", "jane.doe@example.com")
	user.SetName("Jane Smith")
	fmt.Println(user.FullContact())
}
