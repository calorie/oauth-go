package repository

import (
	"github.com/calorie/oauth-go/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepositoty(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUser(email string, password string) (*domain.User, *gorm.DB) {
	user := domain.User{}
	db := r.db.Where("email = ? AND password = ?", email, password).First(&user)
	return &user, db
}
