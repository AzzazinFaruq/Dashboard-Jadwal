package database

import (
	"backend/models"
	"context"
	"log"
)

func Migrate() {
	ctx := context.Background()

	_, err := DB.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table users: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.Kelas)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table kelas: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.MataKuliah)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table mata_kuliah: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.Ruangan)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table ruangan: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.Semester)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table semester: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.Mahasiswa)(nil)).
		ForeignKey(`(user_id) REFERENCES users(id)`).
		ForeignKey(`(kelas_id) REFERENCES kelas(id)`).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table mahasiswa: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.PJMatkul)(nil)).
		ForeignKey(`(user_id) REFERENCES users(id)`).
		ForeignKey(`(mata_kuliah_id) REFERENCES mata_kuliah(id)`).
		ForeignKey(`(kelas_id) REFERENCES kelas(id)`).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table pj_matkul: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.JadwalDefault)(nil)).
		ForeignKey(`(kelas_id) REFERENCES kelas(id)`).
		ForeignKey(`(mata_kuliah_id) REFERENCES mata_kuliah(id)`).
		ForeignKey(`(ruangan_id) REFERENCES ruangan(id)`).
		ForeignKey(`(semester_id) REFERENCES semester(id)`).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table jadwal_default: %v", err)
	}

	_, err = DB.NewCreateTable().
		Model((*models.OverrideJadwal)(nil)).
		ForeignKey(`(jadwal_id) REFERENCES jadwal_default(id)`).
		ForeignKey(`(ruangan_baru) REFERENCES ruangan(id)`).
		ForeignKey(`(dibuat_oleh) REFERENCES users(id)`).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Printf("Error migrating table override_jadwal: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
