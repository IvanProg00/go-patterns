package facade

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) Sell(u User, product string) error {
	if err := u.Card.CheckBalance(); err != nil {
		return err
	}

	for _, prod := range s.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > u.GetBalance() {
			return ErrNotEnoughMoneyOnBalance
		}

		return nil
	}

	return ErrProductNotFound
}
