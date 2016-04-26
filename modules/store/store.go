package store

import (
	"github.com/asdine/storm"
	"gitlab.intelstat.co/timecraft/modules/configuration"
)

var db *storm.DB

// GetDB returns a fetched storm database instance and closes
// the storm database initializer by defer method
func GetDB() *storm.DB {
	defer db.Close()
	if db == nil {
		InitializeStore()
	}
	return db
}

// InitStore creates a storm database instance
func InitStore() {
	// StormDB opens with a default codec of JSON
	db, _ = storm.Open(configuration.Configs.Store)
}
