package pagination

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

var (
	MinPage      = 1
	DefaultLimit = 0
)

type OrderDirection string

const (
	OrderDesc OrderDirection = "DESC"
	OrderAsc  OrderDirection = "ASC"
)

type Order struct {
	Field          string
	OrderDirection OrderDirection
}

func (o Order) String() string {
	return fmt.Sprintf("%v %v", o.Field, o.OrderDirection)
}

type Pagination struct {
	Page       int     `json:"page" pagination:"page"`
	Limit      int     `json:"limit" pagination:"limit"`
	Total      int64   `json:"total"`
	TotalPage  int64   `json:"total_page"`
	Order      string  `json:"order" pagination:"order"`
	InnerOrder []Order `json:"-"`
}

func (p *Pagination) parseOrder() {
	if len(p.Order) == 0 {
		return
	}
	orders := strings.Split(p.Order, ",")
	for _, order := range orders {
		if len(order) < 2 {
			continue
		}
		if order[0] == '-' {
			p.InnerOrder = append(p.InnerOrder, Order{
				Field:          order[1:],
				OrderDirection: OrderDesc,
			})
			continue
		}
		if order[0] == '+' {
			p.InnerOrder = append(p.InnerOrder, Order{
				Field:          order[1:],
				OrderDirection: OrderAsc,
			})
		}
	}
}

func (p *Pagination) Apply(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit
	db = db.Offset(offset).Limit(p.Limit)
	for _, order := range p.InnerOrder {
		db = db.Order(order.String())
	}
	return db
}

func (p *Pagination) Correct() {
	p.parseOrder()
	if p.Limit <= 0 {
		p.Limit = DefaultLimit
	}
	if p.Page < MinPage {
		p.Page = MinPage
	}
}

func (p *Pagination) SetTotal(total int64) {
	p.Total = total
	p.TotalPage = total / int64(p.Limit)
	if total%int64(p.Limit) != 0 {
		p.TotalPage++
	}
}
