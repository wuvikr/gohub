package migrate

import (
	"gohub/pkg/database"

	"gorm.io/gorm"
)

// Migrator 迁移器
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 迁移表
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建迁移器
func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}

	migrator.createMigrationsTable()

	return migrator
}

// createMigrationsTable 创建 migrations 表
func (m *Migrator) createMigrationsTable() {
	migration := &Migration{}

	// 如果没有 migrations 表，则创建
	if !m.Migrator.HasTable(migration) {
		m.Migrator.CreateTable(&migration)
	}
}
