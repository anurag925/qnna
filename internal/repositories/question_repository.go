// internal/repositories/question_repository.go

package repositories

import (
	"github.com/anurag925/qnna/internal/models"
	"github.com/uptrace/bun"
)

type QuestionRepository struct {
	DB *baseRepo[models.Question]
}

func NewQuestionRepository(db bun.IDB) *QuestionRepository {
	return &QuestionRepository{newBaseRepo[models.Question](db)}
}
