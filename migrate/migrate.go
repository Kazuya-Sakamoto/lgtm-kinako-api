package main

import (
	"fmt"

	"lgtm-kinako-api/db"
	"lgtm-kinako-api/domain"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func main() {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)

	migrateDB(dbConn)
}

func migrateDB(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "001_create_users",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.User{})
			},
		},
		{
			ID: "002_create_albums_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Album{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Album{})
			},
		},
		{
			ID: "20231111_remove_album_fields",
			Migrate: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&domain.Album{}, "image_by_data") {
					if err := tx.Migrator().DropColumn(&domain.Album{}, "image_by_data"); err != nil {
						return err
					}
				}
				if tx.Migrator().HasColumn(&domain.Album{}, "bydata") {
					return tx.Migrator().DropColumn(&domain.Album{}, "bydata")
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "003_create_tags_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.Tag{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.Tag{})
			},
		},
		{
			ID: "004_album_tags_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.AlbumTag{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&domain.AlbumTag{})
			},
		},
	})

	if err := m.Migrate(); err != nil {
		fmt.Printf("Could not migrate: %v\n", err)
		return
	}
	fmt.Println("👍---Migration-Successfully---👍")
}
