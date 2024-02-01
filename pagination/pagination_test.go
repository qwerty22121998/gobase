package pagination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_String(t *testing.T) {
	order := Order{
		Field:          "id",
		OrderDirection: OrderDesc,
	}
	assert.Equal(t, "id DESC", order.String())
}

func TestPagination(t *testing.T) {
	p := Pagination{
		Page:       0,
		Limit:      -10,
		Total:      0,
		TotalPage:  0,
		Order:      "",
		InnerOrder: nil,
	}
	p.Correct()
	assert.Equal(t, MinPage, p.Page)
	assert.Equal(t, MinLimit, p.Limit)
	assert.Nil(t, p.InnerOrder)
	p.Order = "-id,,abc,+name"
	p.Correct()
	expectOrder := []Order{
		{"id", OrderDesc}, {"name", OrderAsc},
	}
	assert.Equal(t, expectOrder, p.InnerOrder)
}

func TestPagination_SetTotal(t *testing.T) {
	p := Pagination{
		Limit: 50,
	}
	p.Correct()
	p.SetTotal(61)
	assert.Equal(t, int64(61), p.Total)
	assert.Equal(t, int64(2), p.TotalPage)
	p.SetTotal(100)
	assert.Equal(t, int64(100), p.Total)
	assert.Equal(t, int64(2), p.TotalPage)
}
