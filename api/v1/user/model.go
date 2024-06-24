package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
}

var mockUsers = map[string]User{
	"123e4567-e89b-12d3-a456-426614174000": {
		ID:          uuid.Must(uuid.Parse("123e4567-e89b-12d3-a456-426614174000")),
		Username:    "john_doe",
		Email:       "john_doe@example.com",
		DisplayName: "John Doe",
	},
	"223e4567-e89b-12d3-a456-426614174001": {
		ID:          uuid.Must(uuid.Parse("223e4567-e89b-12d3-a456-426614174001")),
		Username:    "jane_smith",
		Email:       "jane_smith@example.com",
		DisplayName: "Jane Smith",
	},
}
