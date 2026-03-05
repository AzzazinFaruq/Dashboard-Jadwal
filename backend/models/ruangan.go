package models

import "github.com/uptrace/bun"

type Ruangan struct {
	bun.BaseModel `bun:"table:ruangan"`

	ID         int64  `bun:",pk,autoincrement" json:"id"`
	NamaRuangan string `bun:"type:varchar(50),notnull" json:"nama_ruangan"`
	Gedung     string `bun:"type:varchar(50),notnull" json:"gedung"`
}

