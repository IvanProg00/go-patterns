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
