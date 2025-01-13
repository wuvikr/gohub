package config

import "gohub/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{

			"perpage": 10, // 默认分页

			"url_query_page":     "page",     // 分页参数
			"url_query_sort":     "sort",     // 排序参数
			"url_query_order":    "order",    // 排序方式参数
			"url_query_per_page": "per_page", // 每页数量参数
		}
	})
}
