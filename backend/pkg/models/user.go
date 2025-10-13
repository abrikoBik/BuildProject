package models

import (
	"bytes"
	"mp1/pkg/consts"
	"mp1/pkg/l10n"
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
	Login       string    `json:"login" groups:"public" gorm:"unique;not null"`
	Email       string    `json:"email" groups:"public" gorm:"unique;not null"`
	PhoneNumber uint      `json:"phoneNumber" groups:"public" gorm:"unique;not null"`
	Name        string    `json:"name" groups:"public" gorm:"not null"`
	Hash        []byte    `gorm:"not null"`
	Salt        []byte    `gorm:"not null"`
	Sessions    []Session `gorm:"foreignKey:UserID"`
	Projects    []Project `gorm:"foreignKey:UserID"`
	Worker      *Worker
}

type UserRequest struct {
	*User
	Password string `json:"password"`
}

type UserResponse struct {
	*User
	ID uint `json:"id" groups:"register"`
}

func (user *User) Get(db *gorm.DB, id uint, preloads ...string) error {
	return util.Preload(db, preloads...).Where("id = ?", id).Take(&user).Error
}

func (user *User) Create(db *gorm.DB, request UserRequest) (err error) {
	if _, err = mail.ParseAddress(request.Email); err != nil {
		return l10n.InvalidEmailError
	}
	user.Email = request.Email
	if len(request.Password) < consts.MinPasswordLength {
		return l10n.InsecurePasswordError
	}
	if err = user.PreparePassword(request.Password); err != nil {
		return
	}
	user.Name = cases.Title(language.Russian).String(request.Name)
	user.Login = request.Login
	user.PhoneNumber = request.PhoneNumber
	return db.Create(&user).Error
}

func (user *User) Patch(db *gorm.DB, request UserRequest) (err error) {
	if len(request.Password) > 0 {
		if len(request.Password) < consts.MinPasswordLength {
			return l10n.InsecurePasswordError
		}
		if err = user.PreparePassword(request.Password); err != nil {
			return
		}
	}
	if len(request.Email) > 0 {
		if _, err = mail.ParseAddress(request.Email); err != nil {
			return l10n.InvalidEmailError
		}
		user.Email = request.Email
	}
	if len(request.Name) > 0 {
		user.Name = cases.Title(language.Russian).String(request.Name)
	}
	if request.PhoneNumber != 0 {
		user.PhoneNumber = request.PhoneNumber
	}
	return db.Save(&user).Error
}

func (user *User) AddProject(db *gorm.DB, project Project) error {
	user.Projects = append(user.Projects, project)
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
	return (bytes.Equal(user.Hash, hash)), nil
}

func (user User) VerifyToken(db *gorm.DB, token jwt.Token) (result bool, err error) {
	if issuer, err := token.Claims.GetIssuer(); err != nil || issuer != util.JWTIssuer {
		return false, err
	}
	if sub, err := token.Claims.GetSubject(); err != nil || sub != user.Login {
		return false, err
	}
	var (
		wg    sync.WaitGroup
		found = make(chan bool)
	)
	wg.Add(len(user.Sessions))
	for _, session := range user.Sessions {
		go func(foundChan chan<- bool, current string) {
			defer wg.Done()
			if current != token.Raw {
				return
			}
			if !result {
				foundChan <- true
				close(foundChan)
			}
		}(found, session.Token)
	}
	go func(foundChan chan<- bool) {
		wg.Wait()
		if result {
			return
		}
		foundChan <- false
		close(foundChan)
	}(found)
	result = <-found
	return
}
