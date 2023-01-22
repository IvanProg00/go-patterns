# Bridge

Bridge is a structural design pattern that lets you split a large class or a set
of closely related classes into two separate hierarchies—abstraction and
implementation—which can be developed independently of each other.

## Example

We need to implement logic so that our printers and computers can work with each
other in any combination.

1. We need to implement interfaces for scanners and computers.

   ```go
   // scanner/scanner.go

   type Scanner interface {
     ScanFile()
   }
   ```

   ```go
   // pc/pc.go

   type PC interface {
     AddScanner(scanner scanner.Scanner)
     Scan()
   }
   ```

2. Now we will create structures for our computers **Linux**, **Mac** and
   **Windows**. We will write code only for Linux, but the same code will be for
   other computers.

   ```go
   // pc/linux.go

   type Linux struct {
     scanner scanner.Scanner
   }

   func (pc *Linux) AddScanner(sc scanner.Scanner) {
     pc.scanner = sc
   }

   func (pc *Linux) Scan() {
     fmt.Println("scan linux")
     pc.scanner.ScanFile()
   }
   ```

3. And all we have to do is write the code for our scanners **HP** and **Expon**.
   We will write code only for **HP**.

   ```go
   // scanner/hp.go

   type HP struct{}

   func (s *HP) ScanFile() {
     fmt.Println("HP scans file")
   }
   ```
