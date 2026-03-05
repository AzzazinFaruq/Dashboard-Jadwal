package models

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type Mahasiswa struct {
	bun.BaseModel `bun:"table:mahasiswa"`

	ID       int64         `bun:",pk,autoincrement" json:"id"`
	UserID   sql.NullInt64 `bun:"type:bigint,nullzero" json:"user_id"`
	NRP      int64         `bun:"type:bigint,notnull,unique" json:"nrp"`
	FullName string        `bun:"type:varchar(100),notnull" json:"full_name"`
	KelasID  int64         `bun:"type:bigint,notnull" json:"kelas_id"`
	Angkatan int           `bun:"type:int,notnull" json:"angkatan"`
	Prodi    string        `bun:"type:varchar(100),notnull" json:"prodi"`

	User  User  `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	Kelas Kelas `bun:"rel:belongs-to,join:kelas_id=id" json:"kelas,omitempty"`
}

