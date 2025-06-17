package auth

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Password    string    `json:"-"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

func (u *User) UpdatePassword(newPassword string) error {
	if newPassword == "" || len(newPassword) < 6 {
		return ErrPasswordTooShort
	}
	u.Password = newPassword
	u.UpdatedAt = time.Now()
	return nil
}
