// internal/models/answer.go

package models

import (
	"context"
	"time"

	"github.com/anurag925/qnna/internal/utils"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Answer struct {
	bun.Model `bun:"table:answers,primary_key:id"`

	ID        uuid.UUID `bun:"id" json:"id"`
	CreatedAt time.Time `bun:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bun:"updatedAt" json:"updatedAt"`

	QuestionID int64  `bun:"questionId,notnull" json:"questionId"`
	Content    string `json:"content"`
	IsCorrect  bool   `json:"isCorrect"`

	Question Question `bun:"rel:has-one,join:questionId=id"`
}

var _ bun.BeforeAppendModelHook = (*Answer)(nil)

func (m *Answer) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		if m.ID == uuid.Nil {
			m.ID = uuid.New()
		}
		m.CreatedAt = utils.Now()
		m.UpdatedAt = utils.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = utils.Now()
	}
	return nil
}
