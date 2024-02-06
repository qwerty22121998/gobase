package preload

import (
	"github.com/qwerty22121998/gobase/test"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

type mockUser struct {
	ID          uint
	Name        string
	MockProfile *mockProfile `gorm:"foreignKey:UserID"`
}

func (mockUser) TableName() string {
	return "mock_user"
}

type mockProfile struct {
	UserID    uint
	Desc      string
	DeletedAt *time.Time
}

func (mockProfile) TableName() string {
	return "mock_profile"
}

func TestPreload_None(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mock_user`")).
		WillReturnRows(mock.NewRows([]string{"id", "name"}).AddRow(1, "vuhk").AddRow(2, "vuho"))
	var res []mockUser

	p := &preload{
		field: "",
		ptype: "",
		args:  nil,
	}

	db = p.Apply(db.Model(mockUser{})).Find(&res)

	assert.NoError(t, db.Error)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPreload(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mock_user`")).
		WillReturnRows(mock.NewRows([]string{"id", "name"}).AddRow(1, "vuhk").AddRow(2, "vuho"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mock_profile` WHERE `mock_profile`.`user_id` IN (?,?)")).
		WithArgs(1, 2).
		WillReturnRows(mock.NewRows([]string{"user_id", "desc"}).AddRow(1, "desc").AddRow(2, "desc"))
	var res []mockUser

	p := Preload("MockProfile")
	db = p.Apply(db.Model(mockUser{})).Find(&res)

	assert.NoError(t, db.Error)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPreloadWithCondition(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mock_user`")).
		WillReturnRows(mock.NewRows([]string{"id", "name"}).AddRow(1, "vuhk").AddRow(2, "vuho"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `mock_profile` WHERE `mock_profile`.`user_id` IN (?,?) AND order = ?")).
		WithArgs(1, 2, 2).
		WillReturnRows(mock.NewRows([]string{"user_id", "desc"}).AddRow(1, "desc"))
	var res []mockUser

	p := Preload("MockProfile", "order = ?", 2)
	db = p.Apply(db.Model(mockUser{})).Find(&res)

	assert.NoError(t, db.Error)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestJoin(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT `mock_user`.`id`,`mock_user`.`name`,`MockProfile`.`user_id` AS `MockProfile__user_id`,`MockProfile`.`desc` AS `MockProfile__desc`,`MockProfile`.`deleted_at` AS `MockProfile__deleted_at` FROM `mock_user` LEFT JOIN `mock_profile` `MockProfile` ON `mock_user`.`id` = `MockProfile`.`user_id`")).
		WillReturnRows(mock.NewRows([]string{"id", "name", "user_id", "desc"}).AddRow(1, "vuhk", 1, "desc").AddRow(2, "vuho", 2, "desc"))
	var res []mockUser

	p := Join("MockProfile")
	db = p.Apply(db.Model(mockUser{})).Find(&res)

	assert.NoError(t, db.Error)
	assert.NoError(t, mock.ExpectationsWereMet())
}
