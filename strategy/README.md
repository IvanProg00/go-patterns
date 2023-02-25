# Strategy

Strategy is a behavioral design pattern that lets you define a family of algorithms,
put each of them into a separate class, and make their objects interchangeable.

## Example

We will create a navigation system that will create a route.

1. Firstly we will create an interface `Strategy`.

   ```go
   // navigator/strategy.go

   type Strategy interface {
     Route(startPoint, endPoint int)
   }
   ```

2. Now we will create a structure `Navigator`. It will contain the `Strategy`.

   ```go
   // navigator/navigator.go

   type Navigator struct {
     Strategy
   }

   func (n *Navigator) SetStrategy(s Strategy) {
     n.Strategy = s
   }
   ```

3. And finally we will create our three types of navigation: road, public transport
   and walk.

   ```go
   // navigator/road.go

   type RoadStrategy struct{}

   func (rs *RoadStrategy) Route(startPoint, endPoint int) {
     avgSpeed := 30
     trafficJam := 2

     total := endPoint - startPoint
     totalTime := total * 40 * trafficJam

     fmt.Printf("road A %d to B %d; avg speed %d; traffic jam %d; total %d; total time %d min\n", startPoint, endPoint,
       avgSpeed, trafficJam, total, totalTime)
   }
   ```

   ```go
   // navigator/public_transport.go

   type PublicTransportStrategy struct{}

   func (ps *PublicTransportStrategy) Route(startPoint, endPoint int) {
     avgSpeed := 40

     total := endPoint - startPoint
     totalTime := total * 40

     fmt.Printf("public transport A %d to B %d; avg speed %d; total %d; total time %d min\n", startPoint, endPoint,
       avgSpeed, total, totalTime)
   }
   ```

   ```go
   // navigator/walk.go
   type WalkStrategy struct{}

   func (ps *WalkStrategy) Route(startPoint, endPoint int) {
     avgSpeed := 4

     total := endPoint - startPoint
     totalTime := total * 60

     fmt.Printf("walk A %d to B %d; avg speed %d; total %d; total time %d min\n", startPoint, endPoint,
       avgSpeed, total, totalTime)
   }
   ```
