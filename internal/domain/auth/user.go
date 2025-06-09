package auth

import (
	"time"
)

type User struct {
	ID          string
	Email       string
	FirstName   string
	LastName    string
	Password    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) ChangePassword(newPassword string) {
	u.Password = newPassword
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateProfile(firstName, lastName, phoneNumber string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.PhoneNumber = phoneNumber
	u.UpdatedAt = time.Now()
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}
