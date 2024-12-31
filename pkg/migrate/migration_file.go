package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

// MigrationFile 代表着单个迁移文件
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

// getMigrationFile 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
