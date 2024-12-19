package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 返回单词的复数形式
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular 返回单词的单数形式
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake 将字符串转换为蛇形命名，如 TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel 将字符串转换为驼峰命名，如 topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel 将字符串转换为小驼峰命名，如 topic_comment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
