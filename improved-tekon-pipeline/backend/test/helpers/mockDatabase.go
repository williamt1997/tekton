package helpers

import (
	"homers-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func ConnectTestDB() {
	dsn := "host=localhost user=test password=test dbname=test port=8082 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting To Database: ", err)
	} else {
		TestDB = db
	}
}

func BuildTestDB() {
	ConnectTestDB()

	if err := DropAllTables(); err != nil {
		log.Fatalf("Error dropping tables: %v", err)
	}

	err := TestDB.AutoMigrate(
		&models.PageView{},
	)

	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}

func DropAllTables() error {
	if foreignKeyOn := TestDB.Exec("SET session_replication_role = 'replica';").Error; foreignKeyOn != nil {
		return foreignKeyOn
	}

	var tables []string
	if getTables := TestDB.Raw("SELECT tablename FROM pg_tables WHERE schemaname = 'public';").Scan(&tables).Error; getTables != nil {
		return getTables
	}

	for _, table := range tables {
		if err := TestDB.Migrator().DropTable(table); err != nil {

			return err
		}
	}

	if foreignKeyOff := TestDB.Exec("SET session_replication_role = 'origin';").Error; foreignKeyOff != nil {
		return foreignKeyOff
	}
	return nil
}

func CloseTestDB() {
	if TestDB == nil {
		log.Println("TestDB is already nil, nothing to close.")
		return
	}

	sqlDB, err := TestDB.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve SQL DB from TestDB: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Failed to close TestDB connection: %v", err)
	} else {
		log.Println("TestDB connection closed successfully.")
	}

	TestDB = nil
}
