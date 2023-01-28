package facade

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) CheckBalance(cardNumber string) error {
	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}

		if card.Balance <= 0 {
			return ErrNoMoneyOnBalance
		}

		return nil
	}

	return ErrCardNotFound
}
