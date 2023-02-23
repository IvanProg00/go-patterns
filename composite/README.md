# Composite

Composite is a structural design pattern that lets you compose objects into tree
structures and then work with these structures as if they were individual objects.

## Example

We will create a PC. The PC will contain a central processor, a graphics processor
and a motherboard. The client will be able to find any element of our system.

1. We will create the interface of each component of the computer.

   ```go
   // computer/computer.go

   type Component interface {
     Search(name string)
     GetName() string
   }
   ```

2. Let's implement our GPU and CPU.

   ```go
   // computer/gpu.go

   type GPU struct {
     Name        string
     Description string
   }

   func (g *GPU) Search(name string) {
     if g.Name == name {
       fmt.Printf("Component %q found: %q\n", g.Name, g.Description)
     }
   }

   func (g *GPU) GetName() string {
     return g.Name
   }
   ```

   ```go
   // computer/cpu.go

   type CPU struct {
     Name        string
     Description string
   }

   func (c *CPU) Search(name string) {
     if c.Name == name {
       fmt.Printf("Component %q found: %q\n", c.Name, c.Description)
     }
   }

   func (c *CPU) GetName() string {
     return c.Name
   }
   ```

3. And now we will create structures for the PC and the motherboard. Which will
   contain other components inside.

   ```go
   // computer/motherboard.go

   type Motherboard struct {
     Name        string
     Description string
     Components  []Component
   }

   func (mb *Motherboard) Search(name string) {
     if mb.Name == name {
       fmt.Printf("Component %q found: %q\n", mb.Name, mb.Description)
       return
     }

     for _, comp := range mb.Components {
       comp.Search(name)
     }
   }

   func (mb *Motherboard) GetName() string {
     return mb.Name
   }
   ```

   ```go
   // computer/pc.go

   type PC struct {
     Name        string
     Description string
     Components  []Component
   }

   func (pc *PC) Search(name string) {
     if pc.Name == name {
       fmt.Printf("Component %q found: %q\n", pc.Name, pc.Description)
       return
     }

     for _, comp := range pc.Components {
       comp.Search(name)
     }
   }

   func (pc *PC) GetName() string {
     return pc.Name
   }
   ```
