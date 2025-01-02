package seed

import "gorm.io/gorm"

// 存放所有 Seeder
var seeders []Seeder

// 顺序存放 Seeder 的名称
// 用于控制 Seeder 的执行顺序
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应 database/seeds 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 新增 Seeder到 seeders 切片中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}
