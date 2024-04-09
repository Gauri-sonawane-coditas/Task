package main

import "fmt"

// UserService represents a service responsible for user management
type UserService struct{}

// SaveUser saves user data to a database
func (s *UserService) SaveUser(name string) {
	// Simulate saving user to a database
	fmt.Printf("Saving user '%s' to the database\n", name)
}

// OrderService represents a service responsible for order management
type OrderService struct {
	UserService *UserService // Tight coupling between OrderService and UserService
}

// PlaceOrder places an order and saves the user associated with it
func (s *OrderService) PlaceOrder(userName string) {
	// Tight coupling: directly using UserService's method
	s.UserService.SaveUser(userName)
	fmt.Println("Placing order...")
}

func main() {
	userService := &UserService{}
	orderService := &OrderService{UserService: userService}

	orderService.PlaceOrder("John Doe")
}
