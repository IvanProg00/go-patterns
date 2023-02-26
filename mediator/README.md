# Mediator

Mediator is a behavioral design pattern that lets you reduce chaotic dependencies
between objects. The pattern restricts direct communications between the objects
and forces them to collaborate only via a mediator object.

## Example

We will create a transport system. It will check when passengers and cargo can be
delivered.

1. Firstly, we will create interfaces of `Vehicle` and `Dispatcher`.

   ```go
   // vehicle/vehicle.go

   type Vehicle interface {
     Arrive()
     Go()
     PermitArrive()
   }

   type Dispatcher interface {
     CanArrive(v Vehicle) bool
     NotifyAboutGo()
   }
   ```

2. Now we will create a station manager. It will think about what can be delivered.

   ```go
   // vehicle/station_manager.go

   type StationManager struct {
     Queue        []Vehicle
     PlatformFree bool
   }

   func NewStationManager() *StationManager {
     return &StationManager{
       PlatformFree: true,
     }
   }

   func (s *StationManager) CanArrive(v Vehicle) bool {
     if s.PlatformFree {
       s.PlatformFree = false
       return true
     }

     s.Queue = append(s.Queue, v)

     return false
   }

   func (s *StationManager) NotifyAboutGo() {
     if !s.PlatformFree {
       s.PlatformFree = true
     }

     if len(s.Queue) > 0 {
       firstTrainInQueue := s.Queue[0]
       s.Queue = s.Queue[1:]

       firstTrainInQueue.PermitArrive()
     }
   }
   ```

3. And finally, we will implement passenger and cargo logic.

   ```go
   // vehicle/passenger.go

   type Passenger struct {
     Dispatcher
   }

   func (p *Passenger) Arrive() {
     if !p.CanArrive(p) {
       fmt.Println("passengers: departure delayed")

       return
     }

     fmt.Println("passengers: take your seats")
   }

   func (p *Passenger) Go() {
     fmt.Println("passengers: departure")
     p.Dispatcher.NotifyAboutGo()
   }

   func (p *Passenger) PermitArrive() {
     fmt.Println("passengers: take your seats")
     p.Arrive()
   }
   ```

   ```go
   // vehicle/cargo.go

   type Cargo struct {
     Dispatcher
   }

   func (c *Cargo) Arrive() {
     if !c.CanArrive(c) {
       fmt.Println("truck: loading is completed, passengers are on the platform")
       return
     }

     fmt.Println("truck: shipped")
   }

   func (c *Cargo) Go() {
     fmt.Println("truck: successfully loaded")
     c.Dispatcher.NotifyAboutGo()
   }

   func (c *Cargo) PermitArrive() {
     fmt.Println("truck: loading cargo")
     c.Arrive()
   }
   ```
