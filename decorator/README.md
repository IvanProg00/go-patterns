# Decorator

Decorator is a structural design pattern that lets you attach new behaviors to
objects by placing these objects inside special wrapper objects that contain the
behaviors.

## Example

We will create **server** and **home computers**. And these computers will have a
wrapper of **base computer**.

1. Create base computer.

   ```go
   // pc/base.go

   type BasePC struct {
     CPU int
   }

   func (pc BasePC) GetPrice() float64 {
     return 10.0
   }
   ```

2. Create **Wrapper** to get price of our base computer and other computers.

   ```go
   // pc/wrapper.go

   type Wrapper interface {
     GetPrice() float64
   }
   ```

3. Create **server** and **home** computers with interface **Wrapper**. Inside
   **Wrapper** can be base computer.

   ```go
   // pc/server.go

   type ServerPC struct {
     Wrapper Wrapper
     CPU     int
     Memory  int
   }

   func (pc ServerPC) GetPrice() float64 {
     return pc.Wrapper.GetPrice() * float64(pc.CPU) * float64(pc.Memory)
   }

   // pc/home.go

   type HomePC struct {
     Wrapper      Wrapper
     CPU          int
     GraphicsCard int
   }

   func (pc HomePC) GetPrice() float64 {
     return pc.Wrapper.GetPrice() * float64(pc.CPU) * float64(pc.GraphicsCard)
   }
   ```
