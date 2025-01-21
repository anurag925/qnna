// internal/models/question.go

package models

import (
	"context"
	"time"

	"github.com/anurag925/qnna/internal/utils"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Question struct {
	bun.Model `bun:"table:questions,primary_key:id"`

	ID        uuid.UUID `bun:"id" json:"id"`
	CreatedAt time.Time `bun:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bun:"updatedAt" json:"updatedAt"`

	Title   string `json:"title"`
	Content string `json:"content"`
}

var _ bun.BeforeAppendModelHook = (*Question)(nil)

func (m *Question) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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
