package query

import (
	"github.com/qwerty22121998/gobase/test"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

type MockModel struct {
	Name string
	Age  int
}

func (MockModel) TableName() string {
	return "model"
}

func TestEqual(t *testing.T) {
	c := Equal("name", "vuhk")
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name = ?")).
		WithArgs("vuhk").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLike(t *testing.T) {
	c := Like("name", "vuhk")
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name LIKE ?")).
		WithArgs("%vuhk%").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStartWith(t *testing.T) {
	c := StartWith("name", "vuhk")
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name LIKE ?")).
		WithArgs("vuhk%").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestEndWith(t *testing.T) {
	c := EndWith("name", "vuhk")
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name LIKE ?")).
		WithArgs("%vuhk").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIn(t *testing.T) {
	c := In("name", "vuhk", "vuhk1")
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name IN (?,?)")).
		WithArgs("vuhk", "vuhk1").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNonZero(t *testing.T) {
	c := NonZero("", Equal("name", "vuhk"))
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model`")).
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	c = NonZero("vuhk", Equal("name", "vuhk"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name = ?")).
		WithArgs("vuhk").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAnd(t *testing.T) {
	c := And(Equal("name", "vuhk"), Equal("age", 20))
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name = ? AND age = ?")).
		WithArgs("vuhk", 20).
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOr(t *testing.T) {
	c := Or(Equal("name", "vuhk"), Equal("name", "vu"))
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE name = ? OR name = ?")).
		WithArgs("vuhk", "vu").
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAndOr(t *testing.T) {
	c := And(Or(Equal("name", "vuhk"), Equal("name", "vu")), Equal("age", 20))
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model` WHERE (name = ? OR name = ?) AND age = ?")).
		WithArgs("vuhk", "vu", 20).
		WillReturnRows(mock.NewRows([]string{"name"}).AddRow("vuhk"))
	var res []MockModel

	assert.NoError(t, db.Model(MockModel{}).Scopes(c).Find(&res).Error)

	assert.NoError(t, mock.ExpectationsWereMet())
}
