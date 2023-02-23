package subscription

import "fmt"

type Subscriber struct {
	Name string
}

func (sub *Subscriber) Update(pubName string) {
	fmt.Printf("sending from publisher %q to subscriber %q\n", pubName, sub.GetName())
}

func (sub *Subscriber) GetName() string {
	return sub.Name
}
