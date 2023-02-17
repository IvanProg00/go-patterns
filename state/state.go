package state

import (
	"fmt"

	toystore "github.com/IvanProg00/go-patterns/state/toystore"
)

func Run() {
	toyStore := toystore.New(1, 10)

	if err := toyStore.RequestToy(); err != nil {
		fmt.Println("ERROR:", err)
	}

	if err := toyStore.GiveMoney(10); err != nil {
		fmt.Println("ERROR:", err)
	}

	if err := toyStore.DispenseToy(); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println()

	if err := toyStore.AddToys(2); err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println()

	if err := toyStore.RequestToy(); err != nil {
		fmt.Println("ERROR:", err)
	}

	if err := toyStore.GiveMoney(10); err != nil {
		fmt.Println("ERROR:", err)
	}

	if err := toyStore.DispenseToy(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
