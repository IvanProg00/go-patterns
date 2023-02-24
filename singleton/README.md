# Singleton

Singleton is a creational design pattern that lets you ensure that a class has
only one instance, while providing a global access point to this instance.

## Example

Now we will connect to the database. But we have to connect only once, and in other
cases return the saved connection.

1. Let's create our database connection function. Which will be able to work in parallel.

   ```go
   // db/db.go

   type Database struct{}

   var (
     db *Database
     mx = &sync.Mutex{}
   )

   func Connect() *Database {
     mx.Lock()
     defer mx.Unlock()

     if db != nil {
       fmt.Println("already connected to database")
     } else {
       fmt.Println("connecting to database")
       db = &Database{}
     }

     return db
   }
   ```
