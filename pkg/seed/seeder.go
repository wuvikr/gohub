package seed

import (
	"gohub/pkg/console"
	"gohub/pkg/database"

	"gorm.io/gorm"
)

// Seeder 对应 database/seeds 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// 存放所有 Seeder
var seeders []Seeder

// 顺序存放 Seeder 的名称
// 用于控制 Seeder 的执行顺序
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

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

func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if sdr.Name == name {
			return sdr
		}
	}
	return Seeder{}
}

func RunAll() {
	// 先根据顺序执行
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	// 再执行没有顺序的
	for _, sdr := range seeders {
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}

}

// 运行指定的 seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if sdr.Name == name {
			console.Warning("Running seeder: " + sdr.Name)
			sdr.Func(database.DB)
			return
		}
	}
}
