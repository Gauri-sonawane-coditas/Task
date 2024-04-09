package main


import "fmt"

// UserService represents a service responsible for user management
type UserService interface {
	SaveUser(name string)
}

// DatabaseUserService represents a UserService implementation that saves users to a database
type DatabaseUserService struct{}

// SaveUser saves user data to a database
func (s *DatabaseUserService) SaveUser(name string) {
	// Simulate saving user to a database
	fmt.Printf("Saving user '%s' to the database\n", name)
}

// OrderService represents a service responsible for order management
type OrderService struct {
	UserService UserService // Loose coupling through interface
}

// PlaceOrder places an order and uses the injected UserService to save the user
func (s *OrderService) PlaceOrder(userName string) {
	// Loose coupling: using UserService interface
	s.UserService.SaveUser(userName)
	fmt.Println("Placing order...")
}

func main() {
	userService := &DatabaseUserService{} // Dependency injected
	orderService := &OrderService{UserService: userService}

	orderService.PlaceOrder("John Doe")
}
