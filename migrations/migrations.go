package migrations

import (
	"billing/db/main_db"
	"billing/entities"
)

// MigrateDataTable ...
func MigrateDataTable() {
	main_db.DB.AutoMigrate(&entities.User{})
}
