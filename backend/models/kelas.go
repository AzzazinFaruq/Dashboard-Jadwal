package models

import "github.com/uptrace/bun"

type Kelas struct {
	bun.BaseModel `bun:"table:kelas"`

	ID        int64  `bun:",pk,autoincrement" json:"id"`
	NamaKelas string `bun:"type:varchar(50),notnull" json:"nama_kelas"`
	Angkatan  int    `bun:"type:int,notnull" json:"angkatan"`
	Prodi     string `bun:"type:varchar(100),notnull" json:"prodi"`
}

