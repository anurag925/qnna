// internal/repositories/user_repository.go

package repositories

import (
	"github.com/anurag925/qnna/internal/models"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	DB *baseRepo[models.User]
}

func NewUserRepository(db bun.IDB) *UserRepository {
	return &UserRepository{newBaseRepo[models.User](db)}
}
