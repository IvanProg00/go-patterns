# State

State is a behavioral design pattern that lets an object alter its behavior when
its internal state changes. It appears as if the object changed its class.

## Example

We will create a toy store. And it will contain 4 type of states: **has toy**,
**toy requested**, **has money** and **no toys**.

1. Firstly, we will create the interface of our state.

   ```go
   // toystore/state.go

   type State interface {
     AddToys(count int) error
     RequestToy() error
     GiveMoney(money int) error
     DispenseToy() error
   }
   ```

2. Now we will create our 4 states.

   ```go
   // toystore/has_toy.go

   type HasToyState struct {
     toyStore *ToyStore
   }

   func (s *HasToyState) AddToys(count int) error {
     fmt.Printf("toys added %d\n", count)
     s.toyStore.IncrementItemCount(count)

     return nil
   }

   func (s *HasToyState) RequestToy() error {
     if s.toyStore.toysCount == 0 {
       s.toyStore.SetState(s.toyStore.noToys)
       return ErrNoToys
     }

     fmt.Println("toy requested")

     s.toyStore.SetState(s.toyStore.toyRequested)

     return nil
   }

   func (s *HasToyState) GiveMoney(money int) error {
     return ErrSelectToyFirst
   }

   func (s *HasToyState) DispenseToy() error {
     return ErrSelectToyFirst
   }
   ```

   ```go
   // toystore/toy_requested.go

   type ToyRequestedState struct {
     toyStore *ToyStore
   }

   func (s *ToyRequestedState) AddToys(count int) error {
     return ErrToyDispenseInProgress
   }

   func (s *ToyRequestedState) RequestToy() error {
     return ErrToyAlreadyRequested
   }

   func (s *ToyRequestedState) GiveMoney(money int) error {
     if money < s.toyStore.toysPrice {
       return ErrReceivedMoneyIsLess
     }

     fmt.Println("money received")

     s.toyStore.SetState(s.toyStore.hasMoney)

     return nil
   }

   func (s *ToyRequestedState) DispenseToy() error {
     return ErrGiveMoneyFirst
   }
   ```

   ```go
   // toystore/has_money.go

   type HasMoneyState struct {
     ToyStore *ToyStore
   }

   func (s *HasMoneyState) AddToys(count int) error {
     return ErrToyDispenseInProgress
   }

   func (s *HasMoneyState) RequestToy() error {
     return ErrToyDispenseInProgress
   }

   func (s *HasMoneyState) GiveMoney(money int) error {
     return ErrNoToys
   }

   func (s *HasMoneyState) DispenseToy() error {
     fmt.Println("dispense toy")
     s.ToyStore.toysCount--

     if s.ToyStore.toysCount == 0 {
       s.ToyStore.SetState(s.ToyStore.noToys)
     } else {
       s.ToyStore.SetState(s.ToyStore.hasToy)
     }

     return nil
   }
   ```

   ```go
   // toystore/no_toys.go

   type NoToysState struct {
     toyStore *ToyStore
   }

   func (s *NoToysState) AddToys(count int) error {
     s.toyStore.IncrementItemCount(count)
     s.toyStore.SetState(s.toyStore.hasToy)

     return nil
   }

   func (s *NoToysState) RequestToy() error {
     return ErrNoToys
   }

   func (s *NoToysState) GiveMoney(money int) error {
     return ErrNoToys
   }

   func (s *NoToysState) DispenseToy() error {
     return ErrNoToys
   }
   ```

3. And finally, we will implement structure for our toy store.

   ```go
   // toystore/toystore.go

   type ToyStore struct {
     hasToy       State
     toyRequested State
     hasMoney     State
     noToys       State
     currentState State

     toysCount int
     toysPrice int
   }

   func New(toysCount, toysPrice int) *ToyStore {
     toyStore := &ToyStore{
       toysCount: toysCount,
       toysPrice: toysPrice,
     }

     toyStore.hasToy = &HasToyState{toyStore: toyStore}
     toyStore.toyRequested = &ToyRequestedState{toyStore: toyStore}
     toyStore.hasMoney = &HasMoneyState{ToyStore: toyStore}
     toyStore.noToys = &NoToysState{toyStore: toyStore}

     toyStore.SetState(toyStore.hasToy)

     return toyStore
   }

   func (t *ToyStore) AddToys(count int) error {
     return t.currentState.AddToys(count)
   }

   func (t *ToyStore) RequestToy() error {
     return t.currentState.RequestToy()
   }

   func (t *ToyStore) GiveMoney(money int) error {
     return t.currentState.GiveMoney(money)
   }

   func (t *ToyStore) DispenseToy() error {
     return t.currentState.DispenseToy()
   }

   func (t *ToyStore) SetState(s State) {
     t.currentState = s
   }

   func (t *ToyStore) IncrementItemCount(count int) {
     fmt.Printf("add toys %d\n", count)

     t.toysCount += count
   }
   ```
