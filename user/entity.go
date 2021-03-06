package user

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	AvatarFileName string
	PasswordHash   string
	Role           string
	Token          int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
