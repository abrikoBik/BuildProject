package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique;not null"`
	FullName    string `json:"fullName" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	PhoneNumber uint   `json:"phoneNumber" gorm:"unique"`
	Sessions    []Session
	Hash        []byte `gorm:"not null"`
	Salt        []byte `gorm:"not null"`

	// Only used in a JSON object.
	Password string `json:"password" gorm:"-"`
}

type UserJSON struct {
	Name        string `json:"name"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber uint   `json:"phoneNumber"`
	ID          uint   `json:"id"`
}

func (u *UserJSON) Copy(user User) {
	u.Name = user.Name
	u.FullName = user.FullName
	u.Email = user.Email
	u.PhoneNumber = user.PhoneNumber
	u.ID = user.ID
}
