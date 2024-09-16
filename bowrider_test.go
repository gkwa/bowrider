package bowrider

import "testing"

func TestNewBoat(t *testing.T) {
   boat := NewBoat("Speedy", 20.5, 8)
   if boat.Name != "Speedy" || boat.Length != 20.5 || boat.Capacity != 8 {
   	t.Errorf("NewBoat() didn't set values correctly")
   }
}

func TestDescription(t *testing.T) {
   boat := NewBoat("Speedy", 20.5, 8)
   expected := "Speedy is a 20.50 ft boat that can carry 8 people"
   if desc := boat.Description(); desc != expected {
   	t.Errorf("Description() = %v, want %v", desc, expected)
   }
}

func TestIsBowrider(t *testing.T) {
   tests := []struct {
   	name     string
   	length   float64
   	expected bool
   }{
   	{"Small", 15, false},
   	{"Bowrider", 20, true},
   	{"Large", 25, false},
   }

   for _, tt := range tests {
   	t.Run(tt.name, func(t *testing.T) {
   		boat := NewBoat(tt.name, tt.length, 5)
   		if got := boat.IsBowrider(); got != tt.expected {
   			t.Errorf("IsBowrider() = %v, want %v", got, tt.expected)
   		}
   	})
   }
}

