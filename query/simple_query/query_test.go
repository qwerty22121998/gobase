package simple_query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	q := New()
	q.Equal("age", 20)
	q.Like("name", "John")
	q.StartWith("address", "Hanoi")
	q.EndWith("nation", "Viet Nam")
	q.Cmp("id", ">", 10)
	q.In("type", "A", "B")
	q.Custom("A = ? OR A = ?", 1, 2)
	assert.Equal(t, q["age = ?"], []any{20})
	assert.Equal(t, q["name LIKE ?"], []any{"%John%"})
	assert.Equal(t, q["address LIKE ?"], []any{"Hanoi%"})
	assert.Equal(t, q["nation LIKE ?"], []any{"%Viet Nam"})
	assert.Equal(t, q["id > ?"], []any{10})
	assert.Equal(t, q["type IN ?"], []any{[]any{"A", "B"}})
	assert.Equal(t, q["A = ? OR A = ?"], []any{1, 2})
}
