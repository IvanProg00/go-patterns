# Factory Method

Factory method is a creational design pattern which solves the problem of
creating product objects without specifying their concrete classes.

## Example

We will create a computer factory. Our factory will create 3 types of computers:
`laptop`, `personal computer` and `server`.

1. Create an _interface_ of computer.

   ```go
   // computer/computer.go

   type Computer interface {
     GetType() string
     GetCPU() int
     GetOS() string
   }
   ```

2. Create types of our computers.

   ```go
   // computer/computer.go

   type Type string

   const (
     LaptopType           Type = "laptop"
     PersonalComputerType Type = "personal computer"
     ServerType           Type = "server"
   )
   ```

3. Implement _struct_ for `laptop`. Other structures will be the same.

   ```go
   // computer/laptop.go

   type Laptop struct {
     Type Type
     OS   string
     CPU  int
   }

   func NewLaptop() Computer {
     return &Laptop{
       Type: LaptopType,
       OS:   "MacOS",
       CPU:  8,
     }
   }

   func (l *Laptop) GetType() string {
     return string(l.Type)
   }

   func (l *Laptop) GetCPU() int {
     return l.CPU
   }

   func (l *Laptop) GetOS() string {
     return l.OS
   }
   ```

4. Implement function to get computer.

   ```go
   // factorymethod.go

   var ErrComputerNotFound = errors.New("computer not found")

   func Get(compType Type) (Computer, error) {
     switch compType {
     case LaptopType:
       return NewLaptop(), nil
     case PersonalComputerType:
       return NewPersonalComputer(), nil
     case ServerType:
       return NewServer(), nil
     default:
       return nil, ErrComputerNotFound
     }
   }
   ```
