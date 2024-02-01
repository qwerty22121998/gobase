package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isZero(t *testing.T) {
	zeroValue := []any{
		0,
		0.0,
		false,
		"",
		[]any{},
		nil,
	}
	nonZero := []any{
		1,
		0.1,
		true,
		"hello",
		[]any{1, ""},
	}
	for _, v := range zeroValue {
		assert.True(t, isZero(v))
	}
	for _, v := range nonZero {
		assert.False(t, isZero(v))
	}
}
