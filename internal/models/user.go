// internal/models/user.go

package models

import (
	"context"
	"time"

	"github.com/anurag925/qnna/internal/utils"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.Model `bun:"table:users,primary_key:id"`

	ID        uuid.UUID `bun:"id" json:"id"`
	CreatedAt time.Time `bun:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bun:"updatedAt" json:"updatedAt"`

	Email    string `json:"email"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (m *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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
