package facade

import "fmt"

func Run() {
	bank := Bank{
		Name: "Bank",
	}
	card1 := Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 := Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 := User{
		Name: "buyer-1",
		Card: &card1,
	}
	user2 := User{
		Name: "buyer-2",
		Card: &card2,
	}
	prod := Product{
		Name:  "Cheese",
		Price: 150,
	}
	shop := Shop{
		Name: "Shop",
		Products: []Product{
			prod,
		},
	}

	bank.Cards = append(bank.Cards, card1, card2)

	fmt.Printf("user %q\n", user1.Name)

	if err := shop.Sell(user1, prod.Name); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("user %q\n", user2.Name)

	if err := shop.Sell(user2, prod.Name); err != nil {
		fmt.Println("error:", err)
		return
	}
}
