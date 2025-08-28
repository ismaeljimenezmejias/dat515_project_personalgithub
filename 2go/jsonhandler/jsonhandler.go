package jsonhandler

import (
	"encoding/json"
	"time"
)

// User representa un usuario en la aplicación
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

// MarshalUser convierte un User a JSON
func MarshalUser(user User) ([]byte, error) {
	return json.Marshal(user)
}

// UnmarshalUser convierte JSON a User
func UnmarshalUser(data []byte) (User, error) {
	var user User
	err := json.Unmarshal(data, &user)
	return user, err
}

// MarshalUsers convierte un slice de User a JSON
func MarshalUsers(users []User) ([]byte, error) {
	return json.Marshal(users)
}

// UnmarshalUsers convierte JSON a slice de User
func UnmarshalUsers(data []byte) ([]User, error) {
	var users []User
	err := json.Unmarshal(data, &users)
	return users, err
}
