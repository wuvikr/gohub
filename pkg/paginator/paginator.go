package paginator

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int    // 当前页
	PerPage     int    // 每页显示条数
	TotalPage   int    // 总页数
	TotalCount  int64  // 总条数
	NextPageURL string // 下一页
	PrevPageURL string // 上一页
}

// Paginator 分页器结构体
type Paginator struct {
	BaseURL    string // 用来拼接 URL
	PerPage    int    // 每页显示条数
	Page       int    // 当前页数
	Offset     int    // 偏移量
	TotalCount int64  // 总条数
	TotalPage  int    // 总页数
	Sort       string // 排序字段
	Order      string // 排序方式
	query      *gorm.DB
	ctx        *gin.Context
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	// 初始化分页器实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)

	// 查询数据
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

// initProperties 初始化分页
func (p *Paginator) initProperties(perPage int, baseURL string) {
	p.PerPage = p.getPerPage(perPage)
	p.BaseURL = p.formatBaseURL(baseURL)

	// 排序参数
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func (p *Paginator) getPerPage(perPage int) int {
	// 优先从 URL 参数中获取 PerPage
	queryPerpage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerpage) > 0 {
		perPage = cast.ToInt(queryPerpage)
	}

	// 没有传参，使用默认值
	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}
	return perPage
}

// formatBaseURL 格式化 URL
func (p *Paginator) formatBaseURL(baseURL string) string {
	// 判断 baseURL 是否包含 ？
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

// getTotalPage 获取总页数
func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

// getTotalCount 获取总条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// getCurrentPage 获取当前页
func (p *Paginator) getCurrentPage() int {
	// 优先从 URL 参数中获取用户请求的页数
	queryPage := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if queryPage <= 0 {
		return 1
	}

	if p.TotalPage == 0 {
		return 0
	}

	// 如果请求页数大于总页数，返回总页数
	if queryPage > p.TotalPage {
		return p.TotalPage
	}
	return queryPage
}

// getPageLink 生成分页链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

// getPrevPageURL 获取上一页
func (p *Paginator) getPrevPageURL() string {
	if p.TotalPage <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}

// getNextPageURL 获取下一页
func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}
