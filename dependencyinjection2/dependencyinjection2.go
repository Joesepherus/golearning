package dependencyinjection2

import "fmt"

// Database represents a dependency
type Database struct {
	Name string
}

// UserService depends on a database
type UserService struct {
	DB *Database
}

// NewUserService is a constructor that injects the database dependency
func NewUserService(db *Database) *UserService {
	return &UserService{DB: db}
}

// FetchUser fetches a user using the injected database
func (u *UserService) FetchUser() {
	fmt.Printf("Fetching user from database: %s\n", u.DB.Name)
}

func Setup() {
	// Create the dependency (database)
	db := &Database{Name: "UserDB"}

	// Inject the dependency into the UserService
	userService := NewUserService(db)

	// Use the service, which relies on the injected database
	userService.FetchUser()
}
