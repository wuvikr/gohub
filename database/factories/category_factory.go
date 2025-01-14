package factories

import (
	"gohub/app/models/category"

	"github.com/go-faker/faker/v4"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	// 设置唯一性
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.Word(),
			Description: faker.Sentence(),
		}
		objs = append(objs, categoryModel)
	}

	return objs
}
