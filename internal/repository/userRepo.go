package repository

import (
	"github.com/OzyKleyton/arautos-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindAll() ([]model.User, error)
	FindByID(id uint) (user *model.User, err error)
	FindByEmail(email string) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id uint) (user *model.User, err error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) FindAll() (users []model.User, err error) {
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) FindByID(id uint) (user *model.User, err error) {
	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) FindByEmail(email string) (user *model.User, err error) {
	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Update(user *model.User) (*model.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) Delete(id uint) (user *model.User, err error) {
	if err := u.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}