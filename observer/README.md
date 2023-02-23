# Observer

Observer is a behavioral design pattern that lets you define a subscription
mechanism to notify multiple objects about any events that happen to the object
they’re observing.

## Example

We will create a service that will contain a publisher. And clients can subscribe
to the publisher.

1. Firstly, we will create interfaces for our publisher (`Subject`) and subscriber
   (`Consumer`).

   ```go
   // subscription/subject.go

   type Subject interface {
     Subscribe(c Consumer)
     Unsubscribe(c Consumer)
     Notify()
   }
   ```

   ```go
   // subscription/consumer.go

   type Consumer interface {
     Update(pubName string)
     GetName() string
   }
   ```

2. We will create our publisher

   ```go
   // subscription/publisher.go

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
   ```

3. And finally, we will create our subscriber.

   ```go
   // subscription/subscriber.go

   type Subscriber struct {
     Name string
   }

   func (sub *Subscriber) Update(pubName string) {
     fmt.Printf("sending from publisher %q to subscriber %q\n", pubName, sub.GetName())
   }

   func (sub *Subscriber) GetName() string {
     return sub.Name
   }
   ```
