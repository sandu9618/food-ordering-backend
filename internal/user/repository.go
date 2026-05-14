package user

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
