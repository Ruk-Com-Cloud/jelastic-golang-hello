package seeder

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Seeder interface {
	GetName() string
	Run(ctx context.Context, db *gorm.DB) error
	Rollback(ctx context.Context, db *gorm.DB) error
}

type SeederManager struct {
	db      *gorm.DB
	seeders []Seeder
}

func NewSeederManager(db *gorm.DB) *SeederManager {
	return &SeederManager{
		db:      db,
		seeders: make([]Seeder, 0),
	}
}

func (sm *SeederManager) Register(seeder Seeder) {
	sm.seeders = append(sm.seeders, seeder)
}

func (sm *SeederManager) RunAll(ctx context.Context) error {
	log.Println("Starting database seeding...")
	
	for _, seeder := range sm.seeders {
		log.Printf("Running seeder: %s", seeder.GetName())
		
		if err := seeder.Run(ctx, sm.db); err != nil {
			return fmt.Errorf("failed to run seeder %s: %w", seeder.GetName(), err)
		}
		
		log.Printf("Completed seeder: %s", seeder.GetName())
	}
	
	log.Println("Database seeding completed successfully")
	return nil
}

func (sm *SeederManager) RollbackAll(ctx context.Context) error {
	log.Println("Starting database seeding rollback...")
	
	// Run rollbacks in reverse order
	for i := len(sm.seeders) - 1; i >= 0; i-- {
		seeder := sm.seeders[i]
		log.Printf("Rolling back seeder: %s", seeder.GetName())
		
		if err := seeder.Rollback(ctx, sm.db); err != nil {
			return fmt.Errorf("failed to rollback seeder %s: %w", seeder.GetName(), err)
		}
		
		log.Printf("Completed rollback: %s", seeder.GetName())
	}
	
	log.Println("Database seeding rollback completed successfully")
	return nil
}

func (sm *SeederManager) RunSeeder(ctx context.Context, name string) error {
	for _, seeder := range sm.seeders {
		if seeder.GetName() == name {
			log.Printf("Running specific seeder: %s", name)
			
			if err := seeder.Run(ctx, sm.db); err != nil {
				return fmt.Errorf("failed to run seeder %s: %w", name, err)
			}
			
			log.Printf("Completed seeder: %s", name)
			return nil
		}
	}
	
	return fmt.Errorf("seeder %s not found", name)
}

func (sm *SeederManager) ListSeeders() []string {
	names := make([]string, len(sm.seeders))
	for i, seeder := range sm.seeders {
		names[i] = seeder.GetName()
	}
	return names
}