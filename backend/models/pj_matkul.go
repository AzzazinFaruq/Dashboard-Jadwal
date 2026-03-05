package models

import "github.com/uptrace/bun"

type PJMatkul struct {
	bun.BaseModel `bun:"table:pj_matkul"`

	ID          int64 `bun:",pk,autoincrement" json:"id"`
	UserID      int64 `bun:"type:bigint,notnull" json:"user_id"`
	MataKuliahID int64 `bun:"type:bigint,notnull" json:"mata_kuliah_id"`
	KelasID     int64 `bun:"type:bigint,notnull" json:"kelas_id"`

	User      User      `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	MataKuliah MataKuliah `bun:"rel:belongs-to,join:mata_kuliah_id=id" json:"mata_kuliah,omitempty"`
	Kelas     Kelas     `bun:"rel:belongs-to,join:kelas_id=id" json:"kelas,omitempty"`
}

