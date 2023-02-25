# Snapshot

Snapshot is a behavioral design pattern that lets you save and restore the previous
state of an object without revealing the details of its implementation.

## Example

We will change our state and restore to the previous one.

1. Let's create the structure `Snapshot`, that will save the state.

   ```go
   // snapshot.go

   type Snapshot struct {
     state string
   }

   func (s *Snapshot) GetSavedState() string {
     return s.state
   }
   ```

2. Now we will create a structure, that will save all versions of the state.

   ```go
   // guardian.go

   type Guardian struct {
     Items []*Snapshot
   }

   func (g *Guardian) Add(s *Snapshot) {
     g.Items = append(g.Items, s)
   }

   func (g *Guardian) Get(index int) *Snapshot {
     return g.Items[index]
   }
   ```

3. And finally we will create implementation of creating, restoring and modifying
   state.

   ```go
   // creator.go

   type Creator struct {
     State string
   }

   func (c *Creator) Create() *Snapshot {
     return &Snapshot{
       state: c.State,
     }
   }

   func (c *Creator) Restore(s *Snapshot) {
     c.State = s.GetSavedState()
   }

   func (c *Creator) SetState(state string) {
     c.State = state
   }

   func (c *Creator) GetState() string {
     return c.State
   }
   ```
