package observer

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/observer/subscription"
)

func Run() {
	sub1 := &subscription.Subscriber{Name: "sub-1"}
	sub2 := &subscription.Subscriber{Name: "sub-2"}
	sub3 := &subscription.Subscriber{Name: "sub-3"}
	sub4 := &subscription.Subscriber{Name: "sub-4"}

	channel := subscription.Pubslisher{
		Name:      "Publisher channel",
		Consumers: map[string]subscription.Consumer{},
	}

	channel.Subscribe(sub1)
	channel.Subscribe(sub2)
	channel.Subscribe(sub3)
	channel.Subscribe(sub4)

	channel.Notify()

	fmt.Printf("unsubscribe %q\n", sub3.GetName())

	channel.Unsubscribe(sub3)

	channel.Notify()
}
