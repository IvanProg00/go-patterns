package subscription

type Pubslisher struct {
	Consumers map[string]Consumer
	Name      string
}

func (pub *Pubslisher) Subscribe(c Consumer) {
	pub.Consumers[c.GetName()] = c
}

func (pub *Pubslisher) Unsubscribe(c Consumer) {
	delete(pub.Consumers, c.GetName())
}

func (pub *Pubslisher) Notify() {
	for _, c := range pub.Consumers {
		c.Update(pub.Name)
	}
}
