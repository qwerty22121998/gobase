package api

import (
	"encoding/json"
	"github.com/qwerty22121998/gobase/pagination"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {
	r := Success("hello world", "OK").WithDebug("debug")
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `{"meta":{"code":200,"message":"OK","debug":"debug"},"data":"hello world"}`
	assert.JSONEq(t, expected, string(j))
}

func TestSuccessArray(t *testing.T) {
	r := Success([]string{}, "OK")
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `{"meta":{"code":200,"message":"OK"},"data":[]}`
	assert.JSONEq(t, expected, string(j))
}

func TestInfo(t *testing.T) {
	r := Info("OK")
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `{"meta":{"code":200,"message":"OK"}}`
	assert.JSONEq(t, expected, string(j))
}

func TestSuccessPagination(t *testing.T) {
	p := &pagination.Pagination{
		Total:     1,
		Limit:     20,
		TotalPage: 1,
		Page:      1,
	}
	r := SuccessPagination[string]("hello world", p, "OK").WithDebug("debug")
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `
{
  "data": "hello world",
  "meta": {
    "code": 200,
    "debug": "debug",
    "limit": 20,
    "message": "OK",
    "order": "",
    "page": 1,
    "total": 1,
    "total_page": 1
  }
}`
	assert.JSONEq(t, expected, string(j))
}

func TestBadRequest(t *testing.T) {
	r := BadRequest("Bad Request")
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `{"meta":{"code":400,"message":"Bad Request"}}`
	assert.JSONEq(t, expected, string(j))
}

func TestInternalError(t *testing.T) {
	r := InternalError(assert.AnError)
	j, err := json.Marshal(r)
	assert.NoError(t, err)
	expected := `{"meta":{"code":500,"message":"assert.AnError general error for testing"}}`
	assert.JSONEq(t, expected, string(j))
}
