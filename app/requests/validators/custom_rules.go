package validators

import (
	"errors"
	"fmt"
	"gohub/pkg/database"
	"strings"

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
}
