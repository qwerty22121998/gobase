package pagination

import (
	"github.com/qwerty22121998/gobase/test"
	"github.com/stretchr/testify/assert"
	"regexp"
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
	assert.Equal(t, DefaultLimit, p.Limit)
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

func TestPagination_Apply(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	p := Pagination{
		Page:  2,
		Limit: 10,
		Order: "id,-name",
	}
	p.Correct()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model_a` WHERE `model_a`.`deleted_at` IS NULL ORDER BY name DESC LIMIT 10 OFFSET 10")).
		WillReturnRows(mock.NewRows([]string{"id", "name"}).AddRow(1, "vuhk").AddRow(2, "vuho"))

	var res []test.ModelA
	assert.NoError(t, p.Apply(db.Model(test.ModelA{})).Find(&res).Error)
	assert.NoError(t, mock.ExpectationsWereMet())

}
