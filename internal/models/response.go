// internal/models/response.go

package models

import (
	"context"
	"time"

	"github.com/anurag925/qnna/internal/utils"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Response struct {
	bun.Model `bun:"table:responses,primary_key:id"`

	ID        uuid.UUID `bun:"id" json:"id"`
	CreatedAt time.Time `bun:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bun:"updatedAt" json:"updatedAt"`

	UserID     int64  `bun:"userId,notnull" json:"userId"`
	QuestionID int64  `bun:"questionId,notnull" json:"questionId"`
	AnswerID   int64  `bun:"answerId" json:"answerId"`
	Content    string `json:"content"`

	User     User     `bun:"rel:has-one,join:userId=id"`
	Question Question `bun:"rel:has-one,join:questionId=id"`
	Answer   Answer   `bun:"rel:has-one,join:answerId=id"`
}

var _ bun.BeforeAppendModelHook = (*Response)(nil)

func (m *Response) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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
