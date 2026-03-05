package models

import "github.com/uptrace/bun"

type Semester struct {
	bun.BaseModel `bun:"table:semester"`

	ID    int64  `bun:",pk,autoincrement" json:"id"`
	Nama  string `bun:"type:varchar(50),notnull" json:"nama"`
	Tahun int    `bun:"type:int,notnull" json:"tahun"`
	Aktif bool   `bun:"type:boolean,notnull,default:false" json:"aktif"`
}

