# Facade

Facade is a structural design pattern that provides a simplified interface to a
library, a framework, or any other complex set of classes.

## Example

We will create a banking system. Our app will have bank, payment system with
cards, shop, products and users.

1. Let's create structures for Bank, Card, User, Shop and Product.

   ```go
   // bank.go

   type Bank struct {
     Name  string
     Cards []Card
   }
   ```

   ```go
   // card.go

   type Card struct {
     Bank    *Bank
     Name    string
     Balance float64
   }
   ```

   ```go
   // user.go

   type User struct {
     Card *Card
     Name string
   }
   ```

   ```go
   // shop.go

   type Shop struct {
     Name     string
     Products []Product
   }
   ```

   ```go
   // product.go

   type Product struct {
     Name  string
     Price float64
   }
   ```

2. Our shop must sell products, so let's implement it. Firstly, we have to check
   if we have money on our card.

   ```go
   // user.go

   func (u User) GetBalance() float64 {
     return u.Card.Balance
   }
   ```

   ```go
   // bank.go

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
   ```

   ```go
   // card.go

   func (c *Card) CheckBalance() error {
     return c.Bank.CheckBalance(c.Name)
   }
   ```

3. And now we will create a method that will sell the product to the customer.

   ```go
   // shop.go

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
   ```
