# Iterator

Iterator is a behavioral design pattern that lets you traverse elements of a collection
without exposing its underlying representation (list, stack, tree, etc.).

## Example

We will create a travel guide with entries on the routes. And we will get information
about each route.

1. First, we will create our `Iterator` interface. This interface will check if
   the next element exists and receive the element.

   ```go
   // route/route.go

   type Iterator interface {
     HasNext() bool
     GetNext() *Route
   }
   ```

2. Now we need our structure with information about a route.

   ```go
   // route/route.go

   type Route struct {
     Name       string
     TravelTime int
   }
   ```

3. And now, we will use our interface to iterate to the routes.

   ```go
   // route/route.go

   type Routes struct {
     Routes []Route
     index  int
   }

   func (r *Routes) HasNext() bool {
     return r.index < len(r.Routes)
   }

   func (r *Routes) GetNext() *Route {
     if r.HasNext() {
       val := &r.Routes[r.index]
       r.index++

       return val
     }

     return nil
   }
   ```
