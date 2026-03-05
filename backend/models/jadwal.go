package models

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type JadwalDefault struct {
	bun.BaseModel `bun:"table:jadwal_default"`

	ID          int64     `bun:",pk,autoincrement" json:"id"`
	KelasID     int64     `bun:"type:bigint,notnull" json:"kelas_id"`
	MataKuliahID int64    `bun:"type:bigint,notnull" json:"mata_kuliah_id"`
	RuanganID   int64     `bun:"type:bigint,notnull" json:"ruangan_id"`
	SemesterID  int64     `bun:"type:bigint,notnull" json:"semester_id"`
	Hari        string    `bun:"type:enum('senin','selasa','rabu','kamis','jumat','sabtu'),notnull" json:"hari"`
	JamMulai    time.Time `bun:"type:time,notnull" json:"jam_mulai"`
	JamSelesai  time.Time `bun:"type:time,notnull" json:"jam_selesai"`

	Kelas      Kelas     `bun:"rel:belongs-to,join:kelas_id=id" json:"kelas,omitempty"`
	MataKuliah MataKuliah `bun:"rel:belongs-to,join:mata_kuliah_id=id" json:"mata_kuliah,omitempty"`
	Ruangan    Ruangan   `bun:"rel:belongs-to,join:ruangan_id=id" json:"ruangan,omitempty"`
	Semester   Semester  `bun:"rel:belongs-to,join:semester_id=id" json:"semester,omitempty"`
}

type OverrideJadwal struct {
	bun.BaseModel `bun:"table:override_jadwal"`

	ID            int64          `bun:",pk,autoincrement" json:"id"`
	JadwalID      int64          `bun:"type:bigint,notnull" json:"jadwal_id"`
	Tanggal       time.Time      `bun:"type:date,notnull" json:"tanggal"`
	RuanganBaru   sql.NullInt64  `bun:"type:bigint,nullzero" json:"ruangan_baru"`
	JamMulaiBaru  sql.NullTime   `bun:"type:time,nullzero" json:"jam_mulai_baru"`
	JamSelesaiBaru sql.NullTime  `bun:"type:time,nullzero" json:"jam_selesai_baru"`
	Status        string         `bun:"type:enum('normal','pindah','online','batal'),notnull,default:'normal'" json:"status"`
	Alasan        sql.NullString `bun:"type:text,nullzero" json:"alasan"`
	DibuatOleh    int64          `bun:"type:bigint,notnull" json:"dibuat_oleh"`
	CreatedAt     time.Time      `bun:"type:timestamp,nullzero,notnull,default:current_timestamp" json:"created_at"`

	JadwalDefault JadwalDefault `bun:"rel:belongs-to,join:jadwal_id=id" json:"jadwal_default,omitempty"`
	Ruangan       Ruangan       `bun:"rel:belongs-to,join:ruangan_baru=id" json:"ruangan,omitempty"`
	User          User          `bun:"rel:belongs-to,join:dibuat_oleh=id" json:"user,omitempty"`
}