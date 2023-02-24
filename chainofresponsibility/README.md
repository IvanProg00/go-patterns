# Chain of Responsibility

Chain of Responsibility is a behavioral design pattern that lets you pass requests
along a chain of handlers. Upon receiving a request, each handler decides either
to process the request or to pass it to the next handler in the chain.

## Example

We will create a data transfer service. The service will get, update and save data.

1. Let's create interface of the service.

   ```go
   // data/data.go

   type Service interface {
     Execute(data *Data)
     SetNext(svc Service)
   }
   ```

2. We need a `Data` structure that will contain information about the data.

   ```go
   // data/data.go

   type Data struct {
     GetSource    bool
     UpdateSource bool
   }
   ```

3. Now we will create a device service. We will work on data if
   `Data.GetSource == false`

   ```go
   // data/device.go

   type Device struct {
     Next Service
     Name string
   }

   func (d *Device) Execute(data *Data) {
     if data.GetSource {
       fmt.Printf("data from device %q already received\n", d.Name)
       d.Next.Execute(data)

       return
     }

     fmt.Printf("get data from device %q\n", d.Name)

     data.GetSource = true
     d.Next.Execute(data)
   }

   func (d *Device) SetNext(svc Service) {
     d.Next = svc
   }

   ```

4. The same we will do with `UpdateService` and `SaveService`.

   ```go
   // data/update_service.go

   type UpdateService struct {
     Next Service
     Name string
   }

   func (u *UpdateService) Execute(data *Data) {
     if data.UpdateSource {
       fmt.Printf("update data from service %q already received\n", u.Name)
       u.Next.Execute(data)

       return
     }

     fmt.Printf("update data from service %q\n", u.Name)

     data.UpdateSource = true
     u.Next.Execute(data)
   }

   func (u *UpdateService) SetNext(svc Service) {
     u.Next = svc
   }
   ```

   ```go
   // data/save_service.go

   type SaveService struct {
     Next Service
   }

   func (s *SaveService) Execute(data *Data) {
     if !data.UpdateSource {
       fmt.Println("data not updated")
       return
     }

     fmt.Println("data saved")
   }

   func (s *SaveService) SetNext(svc Service) {
     s.Next = svc
   }
   ```
