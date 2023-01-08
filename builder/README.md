# Builder

Builder is a creational design pattern, which allows constructing complex
objects step by step.

## Example

Let's create a factory that will build computers.

1. Create a computer structure.

   ```go
   // computer/computer.go


   type Computer struct {
     Brand       string
     Core        int
     Memory      int
     GraphicCard int
   }

   func (c *Computer) Print() {
     fmt.Printf("%s core: [%d] mem: [%d] graphic card[%d]\n", c.Brand, c.Core, c.Memory, c.GraphicCard)
   }
   ```

2. Create builder interface.

   ```go
   // computer/builder.go

   type Builder interface {
     SetBrand()
     SetCore()
     SetMemory()
     SetGraphicCard()
     GetComputer() Computer
   }
   ```

3. Create a computer constructor for **Asus**.

   ```go
   // computer/asus.go

   type Asus struct {
     Brand       string
     Core        int
     Memory      int
     GraphicCard int
   }

   func (c *Asus) SetBrand() {
     c.Brand = "Asus"
   }

   func (c *Asus) SetCore() {
     c.Core = 4
   }

   func (c *Asus) SetMemory() {
     c.Memory = 6
   }

   func (c *Asus) SetGraphicCard() {
     c.GraphicCard = 1
   }

   func (c *Asus) GetComputer() Computer {
     return Computer{
       Brand:       c.Brand,
       Core:        c.Core,
       Memory:      c.Memory,
       GraphicCard: c.GraphicCard,
     }
   }
   ```

4. Create builder getter.

   ```go
   // computer/builder.go

   type Type int

   const (
     LenovoType Type = iota
     AsusType
   )

   func GetBuilder(builderType Type) Builder {
     switch builderType {
     case LenovoType:
       return &Lenovo{}
     case AsusType:
       return &Asus{}
     default:
       return nil
     }
   }
   ```

5. Implement factory, that will build our computer.

   ```go
   // computer/factory.go

   type Factory struct {
     b Builder
   }

   func NewFactory(b Builder) *Factory {
     return &Factory{
       b: b,
     }
   }

   func (f *Factory) SetBuilder(b Builder) {
     f.b = b
   }

   func (f *Factory) CreateComputer() Computer {
     f.b.SetBrand()
     f.b.SetCore()
     f.b.SetGraphicCard()
     f.b.SetMemory()

     return f.b.GetComputer()
   }
   ```
