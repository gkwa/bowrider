package main

import (
   "fmt"
   "os"
   "strconv"

   "github.com/yourusername/bowrider"
)

func main() {
   if len(os.Args) != 4 {
   	fmt.Println("Usage: boatapp <boat_name> <length> <capacity>")
   	os.Exit(1)
   }

   name := os.Args[1]
   length, err := strconv.ParseFloat(os.Args[2], 64)
   if err != nil {
   	fmt.Println("Error: Length must be a number")
   	os.Exit(1)
   }

   capacity, err := strconv.Atoi(os.Args[3])
   if err != nil {
   	fmt.Println("Error: Capacity must be an integer")
   	os.Exit(1)
   }

   boat := bowrider.NewBoat(name, length, capacity)
   fmt.Println(boat.Description())
   fmt.Printf("Is it a bowrider? %v\n", boat.IsBowrider())
}
