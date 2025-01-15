package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}

	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:varchar(255);index;default:null"`
		UserID     string `gorm:"type:varchar(20);index;default:null"`
		CategoryID string `gorm:"type:varchar(20);index;default:null"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2025_01_15_112159_add_topics_table", up, down)
}
