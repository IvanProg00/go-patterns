# Proxy

Proxy is a structural design pattern that lets you provide a substitute or placeholder
for another object. A proxy controls access to the original object, allowing you
to perform something either before or after the request gets through to the original
object.

## Example

We will get information about the users. But the users must have enough permissions.
So we will create a proxy, that will check if the user can access to the data.

1. Firstly, we will create the interface `Service`.

   ```go
   // db/service.go

   type Service interface {
     GetData(user string) ([]string, error)
   }
   ```

2. Now we need to create database client to get user data. And it will use the interface
   `Service.

   ```go
   // db/db.go

   type Database struct{}

   func (db *Database) GetData(user string) ([]string, error) {
     return []string{
       "line 1",
       "end line",
     }, nil
   }
   ```

3. And now we will create the proxy, that will check if the client has enough rights.

   ```go
   // db/proxy.go

   type ProxyDatabase struct {
     Users map[string]bool
     DB    *Database
   }

   var ErrInsufficientAccessRights = errors.New("insufficient access rights")

   func (pDB *ProxyDatabase) GetData(user string) ([]string, error) {
     if !pDB.Users[user] {
       return nil, ErrInsufficientAccessRights
     }

     return pDB.DB.GetData(user)
   }
   ```
