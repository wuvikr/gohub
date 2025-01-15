package validators

import (
	"errors"
	"fmt"
	"gohub/pkg/database"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

func init() {
	// 自定义验证规则 not_exists, 用于验证某个字段的值在某个表中必须不存在
	// 用于保证数据库表中记录的唯一性，例如用户表中 email、手机号必须唯一
	// not_exists 参数有两种，一种是两个参数，，一种是三个参数，表示检查数据库表中某个字段的值不存在，并排除 id = ? 的情况
	// not_exists:users,email 表示检查 users 表中 email 字段的该值不存在
	// not_exists:users,email,32 表示检查 users 表中 email 字段的该值不存在，并排除 id = 32 的情况
	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，例如 users
		tableName := rng[0]

		// 第二个参数，字段名，例如 email
		fieldName := rng[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)
		query := database.DB.Table(tableName).Where(fieldName+" = ?", requestValue)

		if exceptID != "" {
			query = query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 如果记录存在则验证失败
		if count != 0 {
			// 如果设置了错误消息
			if message != "" {
				return errors.New(message)
			}

			// 默认错误消息
			return fmt.Errorf("%v 已被占用", requestValue)
		}

		// 验证通过
		return nil
	})

	// max_cn 中文最大长度
	govalidator.AddCustomRule("max_cn", func(field, rule, message string, value any) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果设置了错误消息
			if message != "" {
				return errors.New(message)
			}
			// 默认错误消息
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}

		return nil
	})

	// min_cn 中文最小长度
	govalidator.AddCustomRule("min_cn", func(field, rule, message string, value any) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果设置了错误消息
			if message != "" {
				return errors.New(message)
			}
			// 默认错误消息
			return fmt.Errorf("长度不能少于 %d 个字", l)
		}

		return nil
	})

	// exists 自定义规则, 确保字段的值在数据库表中存在
	govalidator.AddCustomRule("exists", func(field, rule, message string, value any) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		// 第一个参数，表名称，例如 users
		tableName := rng[0]

		// 第二个参数，字段名，例如 email
		fieldName := rng[1]

		reqValue := value.(string)

		// 查询数据库
		var count int64
		database.DB.Table(tableName).Where(fieldName+" = ?", reqValue).Count(&count)

		// 如果记录不存在则验证失败
		if count == 0 {
			// 如果设置了错误消息
			if message != "" {
				return errors.New(message)
			}
			// 默认错误消息
			return fmt.Errorf("%v 不存在", reqValue)
		}

		return nil
	})
}
