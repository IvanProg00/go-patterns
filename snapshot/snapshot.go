package snapshot

import (
	"fmt"
)

type Snapshot struct {
	state string
}

func (s *Snapshot) GetSavedState() string {
	return s.state
}

func Run() {
	storage := Guardian{
		Items: []*Snapshot{},
	}

	creator := Creator{
		State: "A",
	}

	fmt.Printf("creator current state: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("B")
	fmt.Printf("creator current state: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("C")
	fmt.Printf("creator current state: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(1))
	fmt.Printf("restored to state: %s\n", creator.GetState())

	creator.Restore(storage.Get(0))
	fmt.Printf("restored to state: %s\n", creator.GetState())
}
