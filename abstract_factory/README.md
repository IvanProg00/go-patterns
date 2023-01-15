# Abstract Factory

Abstract Factory is a creational design pattern that lets you produce families of
related objects without specifying their concrete classes.

## Example

We will create a factory that will create monitors and computers of companies
**Dell** and **HP**.

1. Let's create a computer interface and its implementation.

   ```go
   // factory/computer.go

   type IComputer interface {
     SetMemory(memory int)
     SetCPU(cpu int)
     GetMemory() int
     GetCPU() int
   }

   type Computer struct {
     Memory int
     CPU    int
   }

   func (c *Computer) SetMemory(memory int) {
     c.Memory = memory
   }

   func (c *Computer) SetCPU(cpu int) {
     c.CPU = cpu
   }

   func (c *Computer) GetMemory() int {
     return c.Memory
   }

   func (c *Computer) GetCPU() int {
     return c.CPU
   }
   ```

2. Now we will create our computer implementation for DELL.

   ```go
   // factory/dell_computer.go

   type DellComputer struct {
     Computer
   }
   ```

3. And now we will create our implementation for the monitor.

   ```go
   // factory/monitor.go

   type IMonitor interface {
     SetSize(size int)
     GetSize() int
   }

   type Monitor struct {
     Size int
   }

   func (m *Monitor) SetSize(size int) {
     m.Size = size
   }

   func (m *Monitor) GetSize() int {
     return m.Size
   }
   ```

   ```go
   // factory/dell_monitor.go

   type DellMonitor struct {
     Monitor
   }
   ```

4. Let's create an implementation for HP.

   ```go
   // factory/hp_computer.go

   type HPComputer struct {
     Computer
   }
   ```

   ```go
   // factory/hp_monitor.go

   type HPMonitor struct {
     Monitor
   }
   ```

5. Now let's create an interface for the factory.

   ```go
   // factory/factory.go

   type ComputerFactory interface {
     GetMonitor() IMonitor
     GetComputer() IComputer
   }
   ```

6. And now we will use factory interface `ComputerFactory` for HP and Dell.

   ```go
   // factory/dell.go

   type Dell struct{}

   func (d *Dell) GetComputer() IComputer {
     return &DellComputer{
       Computer: Computer{
         Memory: 8,
         CPU:    4,
       },
     }
   }

   func (d *Dell) GetMonitor() IMonitor {
     return &DellMonitor{
       Monitor: Monitor{
         Size: 27,
       },
     }
   }
   ```

   ```go
   // factory/hp.go

   type HP struct{}

   func (h *HP) GetMonitor() IMonitor {
     return &HPMonitor{
       Monitor: Monitor{
         Size: 24,
       },
     }
   }

   func (h *HP) GetComputer() IComputer {
     return &HPComputer{
       Computer: Computer{
         Memory: 16,
         CPU:    8,
       },
     }
   }
   ```

7. And finally, we will implement a function that will return the factory for
   the brand we need.

   ```go
   // factory/factory.go


   type Brand string

   const (
     DellBrand Brand = "dell"
     HPBrand   Brand = "hp"
   )

   var ErrBrandNotFound = errors.New("brand not found")

   func GetComputerFactory(brand Brand) (ComputerFactory, error) {
     switch brand {
     case DellBrand:
       return &Dell{}, nil
     case HPBrand:
       return &HP{}, nil
     default:
       return nil, ErrBrandNotFound
     }
   }
   ```
