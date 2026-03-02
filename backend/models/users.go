package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Username  string    `bun:"type:varchar(50),notnull" json:"username"`
	Password  string    `bun:"type:varchar(100),notnull" json:"-"`
	Fullname  string    `bun:"type:varchar(100),notnull" json:"fullname"`
	Role      string    `bun:"type:enum('admin', 'penanggungjawab', 'mahasiswa'),notnull" json:"role"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Clinic Clinic `bun:"rel:belongs-to,join:clinic_id=id" json:"clinic,omitempty"`
}
