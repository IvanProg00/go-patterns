package facade

type User struct {
	Card *Card
	Name string
}

func (u User) GetBalance() float64 {
	return u.Card.Balance
}
