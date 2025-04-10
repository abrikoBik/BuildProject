package models

import (
	"bytes"
	"mp1/pkg/util"
	"net/mail"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string    `json:"login" groups:"public" gorm:"unique;not null"`
	Name     string    `json:"name" groups:"public" gorm:"not null"`
	Email    string    `json:"email" groups:"public" gorm:"unique;not null"`
	Hash     []byte    `gorm:"not null"`
	Salt     []byte    `gorm:"not null"`
	Sessions []Session `gorm:"foreignKey:UserID"`
	Worker   Worker    `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

type UserResponse struct {
	*User
	RegisteredAt string `json:"registeredAt" groups:"public"`
	ID           uint   `json:"id" groups:"register"`
}

func (user *User) Create(db *gorm.DB, request UserRequest) (err error) {
	if _, err = mail.ParseAddress(request.Email); err != nil {
		return util.InvalidEmailError
	}
	user.Email = request.Email
	if len(request.Password) < util.MinPasswordLength {
		return util.InsecurePasswordError
	}
	if err = user.PreparePassword(request.Password); err != nil {
		return
	}
	user.Name = cases.Title(language.Russian).String(request.Name)
	user.Login = request.Login
	return db.Create(&user).Error
}

func (user *User) Patch(db *gorm.DB, request UserRequest) (err error) {
	if len(request.Password) > 0 {
		if len(request.Password) < util.MinPasswordLength {
			return util.InsecurePasswordError
		}
		if err = user.PreparePassword(request.Password); err != nil {
			return
		}
	}
	if len(request.User.Email) > 0 {
		if _, err = mail.ParseAddress(request.User.Email); err != nil {
			return util.InvalidEmailError
		}
		user.Email = request.User.Email
	}
	if len(request.User.Name) > 0 {
		user.Name = cases.Title(language.Russian).String(request.User.Name)
	}
	return db.Save(&user).Error
}

func (user *User) PreparePassword(password string) (err error) {
	if user.Salt, err = util.RandomBytes(16); err != nil {
		return
	}
	user.Hash, err = util.GenerateHash([]byte(password), user.Salt)
	return
}

func (user User) VerifyPassword(password string) (bool, error) {
	hash, err := util.GenerateHash([]byte(password), user.Salt)
	if err != nil {
		return false, err
	}
	return (bytes.Compare(user.Hash, hash) == 0), nil
}

func (user User) VerifyToken(db *gorm.DB, token jwt.Token) (bool, error) {
	if issuer, err := token.Claims.GetIssuer(); err != nil || issuer != util.JWTIssuer {
		return false, err
	}
	if sub, err := token.Claims.GetSubject(); err != nil || sub != user.Login {
		return false, err
	}
	query := db.Model(&Session{}).Where("user_id = ?", user.ID).Select("token")
	rows, err := query.Rows()
	defer rows.Close()
	if err != nil {
		return false, err
	}
	var (
		wg    sync.WaitGroup
		count int64
		found = make(chan bool, 1)
	)
	query.Count(&count)
	wg.Add(int(count))
	for rows.Next() {
		var session Session
		db.ScanRows(rows, &session)
		go func() {
			if session.Token == token.Raw {
				found <- true
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		found <- false
		close(found)
	}()
	return <-found, nil
}
