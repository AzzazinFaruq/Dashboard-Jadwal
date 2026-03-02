package models

import (
	"time"

	"github.com/uptrace/bun"
)

type MataKuliah struct {
	bun.BaseModel `bun:"table:mata_kuliah"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	NamaMataKuliah  string    `bun:"type:varchar(50),notnull" json:"nama_mata_kuliah"`
	SKS  int    `bun:"type:int,notnull" json:"sks"`
	DosenPengajar  string    `bun:"type:varchar(100),notnull" json:"dosen_pengajar"`
	
	
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`

	// Clinic Clinic `bun:"rel:belongs-to,join:clinic_id=id" json:"clinic,omitempty"`
}
