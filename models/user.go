package models

import (
	"errors"
	"time"

	"github.com/BharathKumarRavichandran/cloud99/db"
	"github.com/BharathKumarRavichandran/cloud99/utils"

	"golang.org/x/crypto/bcrypt"

	forms "github.com/BharathKumarRavichandran/cloud99/forms/auth"
)

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type User struct {
	UserID    uint   `gorm:"index:user_id;PRIMARY_KEY"`
	Name      string `gorm:"type:varchar(100);NOT NULL"`
	Email     string `gorm:"type:varchar(100);UNIQUE"`
	Password  string `gorm:"type:varchar(100);NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type UserModel struct{}

func (m UserModel) Login(form forms.LoginForm) (user User, err error) {
	db.DB.Where("email = ?", form.Email).First(&user)
	if user.UserID <= 0 {
		return user, errors.New("User account does not exists!")
	}

	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return user, errors.New("Invalid password")
	}

	return user, nil
}

func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
	db.DB.Where("email = ?", form.Email).First(&user)
	if user.UserID > 0 {
		return user, errors.New("User account already exists!")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		utils.Logger.Errorf("UserEmail %s : %s", form.Email, err.Error())
		panic(err)
	}

	user = User{
		Name:     form.Name,
		Email:    form.Email,
		Password: string(hashedPassword),
	}

	db.DB.Create(&user)

	return user, nil
}
