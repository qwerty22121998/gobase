package base_repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/qwerty22121998/gobase/base_model"
	"github.com/qwerty22121998/gobase/pagination"
	"github.com/qwerty22121998/gobase/query"
	"github.com/qwerty22121998/gobase/test"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{},
		A:     "A",
	}

	repo := New[*test.ModelA](db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `model_a` (`created_at`,`updated_at`,`deleted_at`,`created_by`,`updated_by`,`a`) VALUES (?,?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), "", "", data.A).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	assert.NoError(t, repo.Create(context.Background(), data))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Save(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{
			ID: 1,
		},
		A: "A",
	}

	repo := New[*test.ModelA](db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `model_a` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`created_by`=?,`updated_by`=?,`a`=? WHERE `model_a`.`deleted_at` IS NULL AND `id` = ?")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", data.A, data.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	assert.NoError(t, repo.Save(context.Background(), data))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Delete(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{
			ID: 1,
		},
		A: "A",
	}

	repo := New[*test.ModelA](db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `model_a` SET `deleted_at`=? WHERE `model_a`.`id` = ? AND `model_a`.`deleted_at` IS NULL")).
		WithArgs(sqlmock.AnyArg(), data.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	assert.NoError(t, repo.Delete(context.Background(), data))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_FindByID(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{
			ID: 1,
		},
		A: "A",
	}

	repo := New[*test.ModelA](db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model_a` WHERE `model_a`.`id` = ? AND `model_a`.`deleted_at` IS NULL ORDER BY `model_a`.`id` LIMIT 1")).
		WithArgs(data.ID).
		WillReturnRows(mock.NewRows([]string{"id", "a"}).AddRow(data.ID, data.A))

	result, err := repo.FindByID(context.Background(), data.ID)
	assert.NoError(t, err)
	assert.Equal(t, data, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_FindFirst(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{
			ID: 1,
		},
		A: "A",
	}

	q := query.And(query.Equal("a", "A"))

	repo := New[*test.ModelA](db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model_a` WHERE a = ? AND `model_a`.`deleted_at` IS NULL ORDER BY `model_a`.`id` LIMIT 1")).
		WithArgs("A").
		WillReturnRows(mock.NewRows([]string{"id", "a"}).AddRow(data.ID, data.A))

	result, err := repo.FindFirst(context.Background(), q)
	assert.NoError(t, err)
	assert.Equal(t, data, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_FindMany(t *testing.T) {
	db, mock, err := test.DB()
	assert.NoError(t, err)

	data := &test.ModelA{
		Model: base_model.Model{
			ID: 1,
		},
		A: "A",
	}

	q := query.And(query.Equal("a", "A"))
	p := pagination.Pagination{
		Page:      0,
		Limit:     10,
		Total:     0,
		TotalPage: 0,
		Order:     "-id",
	}

	repo := New[*test.ModelA](db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `model_a` WHERE a = ? AND `model_a`.`deleted_at` IS NULL")).
		WithArgs("A").
		WillReturnRows(mock.NewRows([]string{"count(*)"}).AddRow(1))

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `model_a` WHERE a = ? AND `model_a`.`deleted_at` IS NULL LIMIT 10")).
		WithArgs("A").
		WillReturnRows(mock.NewRows([]string{"id", "a"}).AddRow(data.ID, data.A))

	result, err := repo.FindMany(context.Background(), q, &p)
	assert.NoError(t, err)
	assert.Equal(t, []*test.ModelA{data}, result)
	assert.Equal(t, int64(1), p.Total)
	assert.Equal(t, int64(1), p.TotalPage)
	assert.NoError(t, mock.ExpectationsWereMet())
}
