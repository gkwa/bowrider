package bowrider

import "fmt"

type Boat struct {
   Name     string
   Length   float64
   Capacity int
}

func NewBoat(name string, length float64, capacity int) *Boat {
   return &Boat{
   	Name:     name,
   	Length:   length,
   	Capacity: capacity,
   }
}

func (b *Boat) Description() string {
   return fmt.Sprintf("%s is a %.2f ft boat that can carry %d people", b.Name, b.Length, b.Capacity)
}

func (b *Boat) IsBowrider() bool {
   return b.Length >= 16 && b.Length <= 24
}

