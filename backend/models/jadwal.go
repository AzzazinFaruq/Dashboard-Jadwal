package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Jadwal struct {
	bun.BaseModel `bun:"table:jadwal"`
	
	ID        int64     `bun:",pk,autoincrement" json:"id"`
	MataKuliahID int64     `bun:"type:bigint,notnull" json:"mata_kuliah_id"`
	Ruangan string    `bun:"type:varchar(100),notnull" json:"ruangan"`
	JamMulai string    `bun:"type:varchar(100),notnull" json:"jam_mulai"`
	JamSelesai string    `bun:"type:varchar(100),notnull" json:"jam_selesai"`
	Hari string    `bun:"type:varchar(100),notnull" json:"hari"`
	Kelas string    `bun:"type:varchar(100),notnull" json:"kelas"`
	UserID int64 `bun:"type:bigint,notnull" json:"user_id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`

	MataKuliah MataKuliah `bun:"rel:belongs-to,join:mata_kuliah_id=id" json:"mata_kuliah,omitempty"`
	User User `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
}