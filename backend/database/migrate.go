package database

import (
	"backend/models"
	"context"
	"log"
)

func Migrate() {
	_, err := DB.NewCreateTable().Model((*models.User)(nil)).
	IfNotExists().Exec(context.Background())
	if err != nil {
		log.Printf("Error migrating table users: %v", err)
	}

	_, err = DB.NewCreateTable().Model((*models.MataKuliah)(nil)).IfNotExists().Exec(context.Background())
	if err != nil {
		log.Printf("Error migrating table mata_kuliah: %v", err)
	}

	_, err = DB.NewCreateTable().Model((*models.Jadwal)(nil)).
	ForeignKey(`(mata_kuliah_id) REFERENCES mata_kuliah(id) ON DELETE CASCADE`).
	ForeignKey(`(user_id) REFERENCES users(id) ON DELETE CASCADE`).
	IfNotExists().Exec(context.Background())
	if err != nil {
		log.Printf("Error migrating table jadwal: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
