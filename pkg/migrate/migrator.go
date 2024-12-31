package migrate

import (
	"fmt"
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gohub/pkg/file"
	"os"

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

// Up 执行迁移
func (m *Migrator) Up() {
	// 获取所有的迁移文件
	migrateFiles := m.readAllMigrationFiles()

	// 获取当前批次的值
	batch := m.getBatch()

	// 获取 migrations 表中的所有记录
	migrations := []Migration{}
	m.DB.Find(&migrations)

	// 通过这个参数来判断数据库是否已经最新
	runed := false
	for _, mfile := range migrateFiles {

		if mfile.isNotMigrated(migrations) {
			m.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}
}

// Rollback 回滚上一次迁移
func (m *Migrator) Rollback() {
	// 获取最后一次迁移的记录
	lastMigration := Migration{}
	m.DB.Order("id desc").First(&lastMigration)
	migrations := []Migration{}
	m.DB.Where("batch = ?", lastMigration.Batch).Order("id desc").Find(&migrations)

	// 回滚最后一批次的迁移
	if !m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// rollbackMigrations 回滚迁移
func (m *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 标记是否真的有执行了迁移回退的操作
	runed := false

	for _, _migration := range migrations {
		// 提示信息
		console.Warning(fmt.Sprintf("Rolling back: %s", _migration.Migration))

		// 执行迁移文件的 Down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}

		runed = true

		// 删除迁移记录
		m.DB.Delete(&_migration)
		// 提示已回滚文件信息
		console.Success(fmt.Sprintf("Rolled back: %s", mfile.FileName))
	}
	return runed
}

func (m *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations 目录下的所有文件
	files, err := os.ReadDir(m.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// 去除文件后缀
		fileName := file.FileNameWithoutExtension(f.Name())

		//
		mfile := getMigrationFile(fileName)

		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	return migrateFiles
}

// getBatch 获取当前批次的值
func (m *Migrator) getBatch() int {
	// 默认批次值为 1
	batch := 1

	// 获取 migrations 表中的最大批次值
	lastMigration := Migration{}
	m.DB.Order("id desc").First(&lastMigration)

	// 如果有记录，则批次值加 1
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}

	return batch
}

func (m *Migrator) runUpMigration(mfile MigrationFile, batch int) {

	// 如果 Up 不为 nil，则执行迁移
	if mfile.Up != nil {
		// 提示信息
		console.Warning(fmt.Sprintf("Migrating: %s", mfile.FileName))
		// 执行迁移
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移文件信息
		console.Success(fmt.Sprintf("Migrated: %s", mfile.FileName))
	}

	// 记录迁移记录
	err := m.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

// Reset 回滚所有迁移
func (m *Migrator) Reset() {
	migrations := []Migration{}

	// 倒序获取 migrations 表中的所有记录
	m.DB.Order("id desc").Find(&migrations)

	// 回滚所有迁移
	if !m.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

// Refresh 重置并重新执行迁移
func (m *Migrator) Refresh() {
	m.Reset()
	m.Up()
}
