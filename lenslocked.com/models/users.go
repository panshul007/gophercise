package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound Error returned when resource not found.
	ErrNotFound  = errors.New("models: resource not found")
	ErrInvalidId = errors.New("models: ID provided was invalid")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

type UserService struct {
	db *gorm.DB
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

// Closes the user service database connection.
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset drops the user table and re creates it.
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}

// ByID will lookup user by id provided.
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, otherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// ByEmail will fetch the user by provided email.
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Create will create the provided user and backfill data
// like ID, CreatedAt, etc.
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Update will update the provided the user with all of the data
// provided in the user object.
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delete the user with provided user Id.
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// first will fetch the first record by the provided gorm db condition
// and place it in the dst object.
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}