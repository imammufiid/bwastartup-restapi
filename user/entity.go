package user

import "time"

// create struct for user entity
// merepresentasikan field yang ada di database

type User struct {
	ID                                                          int
	Name, Occupation, Email, PasswordHash, Avatar, Role string
	CreatedAt, UpdatedAt                                        time.Time
}
