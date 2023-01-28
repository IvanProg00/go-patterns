package facade

type Card struct {
	Bank    *Bank
	Name    string
	Balance float64
}

func (c *Card) CheckBalance() error {
	return c.Bank.CheckBalance(c.Name)
}
