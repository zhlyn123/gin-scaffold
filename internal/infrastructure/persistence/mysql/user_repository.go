package mysql

import (
	"context"
	"gin-scaffold/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id uint,
) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
