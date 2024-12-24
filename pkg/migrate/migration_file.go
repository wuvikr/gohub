package migrate

import "gorm.io/gorm"

type migrationFunc func(gorm.Migrator, *gorm.DB)

// MigrationFile 迁移单个文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

// migrationFiles 保存所有的迁移记录
var migrationFiles []MigrationFile

// Add 新增一条迁移记录
func Add(name string, up, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
