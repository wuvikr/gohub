package factories

import (
	"gohub/app/models/topic"

	"github.com/go-faker/faker/v4"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			UserID:     "1",
			CategoryID: "3",
		}
		objs = append(objs, topicModel)
	}

	return objs
}
