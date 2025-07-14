package seeder

import "gorm.io/gorm"

func RegisterAllSeeders(manager *SeederManager) {
	// Register all seeders in dependency order
	manager.Register(NewUserSeeder())
	
	// Add future seeders here in the correct order
	// For example:
	// manager.Register(NewCategorySeeder())
	// manager.Register(NewProductSeeder())
}

func SetupSeeders(db *gorm.DB) *SeederManager {
	manager := NewSeederManager(db)
	RegisterAllSeeders(manager)
	return manager
}