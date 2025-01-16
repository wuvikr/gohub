package factories

import (
	"gohub/app/models/link"

	"github.com/go-faker/faker/v4"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			Name: faker.Word(),
			URL:  faker.URL(),
		}
		objs = append(objs, linkModel)
	}

	return objs
}
