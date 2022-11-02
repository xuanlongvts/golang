package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create a user
func CreateUser(db *gorm.DB, user *Users) error {
	err := db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

// get users
func GetUsers(db *gorm.DB, user *[]Users) (err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}

// get user by id
func GetUser(db *gorm.DB, user *Users, email string) (err error) {
	err = db.Where("email = ?", email).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// update user
func UpdateUser(db *gorm.DB, user *Users, email string) (err error) {
	err = db.Where("email = ?", email).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// delete user
func DeleteUser(db *gorm.DB, user *Users, email string) (int64, error) {
	err := db.Where("email = ?", email).Delete(user).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}
