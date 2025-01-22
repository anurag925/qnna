// internal/repositories/response_repository.go

package repositories

import (
	"github.com/anurag925/qnna/internal/models"
	"github.com/uptrace/bun"
)

type ResponseRepository struct {
	DB *baseRepo[models.Response]
}

func NewResponseRepository(db bun.IDB) *ResponseRepository {
	return &ResponseRepository{newBaseRepo[models.Response](db)}
}
