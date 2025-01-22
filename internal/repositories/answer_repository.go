// internal/repositories/answer_repository.go

package repositories

import (
	"github.com/anurag925/qnna/internal/models"
	"github.com/uptrace/bun"
)

type AnswerRepository struct {
	DB *baseRepo[models.Answer]
}

func NewAnswerRepository(db bun.IDB) *AnswerRepository {
	return &AnswerRepository{newBaseRepo[models.Answer](db)}
}
