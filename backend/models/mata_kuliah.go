package models

import (
	"github.com/uptrace/bun"
)

type MataKuliah struct {
	bun.BaseModel `bun:"table:mata_kuliah"`

	ID     int64  `bun:",pk,autoincrement" json:"id"`
	KodeMK string `bun:"type:varchar(20),notnull" json:"kode_mk"`
	NamaMataKuliah string `bun:"type:varchar(150),notnull,column:nama_mk" json:"nama_mk"`
	SKS    int    `bun:"type:int,notnull" json:"sks"`

}
