package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Username  string    `bun:"type:varchar(50),notnull,unique" json:"username"`
	Password  string    `bun:"type:varchar(255),notnull" json:"-"`
	Role      string    `bun:"type:enum('admin', 'pj', 'mahasiswa'),notnull" json:"role"`
	CreatedAt time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"-" json:"updated_at,omitempty"`
}
