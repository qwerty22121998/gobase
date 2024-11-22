package simple_query

import (
	"fmt"
	"gorm.io/gorm"
)

type Query map[string][]any

func New() Query {
	return make(Query)
}

func (s Query) Apply(db *gorm.DB) *gorm.DB {
	for k, v := range s {
		db = db.Where(k, v...)
	}
	return db
}

func (s Query) Equal(key string, value any) Query {
	s[fmt.Sprintf("%v = ?", key)] = []any{value}
	return s
}

func (s Query) Like(key string, value string) Query {
	s[fmt.Sprintf("%v LIKE ?", key)] = []any{fmt.Sprintf("%%%v%%", value)}
	return s
}

func (s Query) StartWith(key string, value string) Query {
	s[fmt.Sprintf("%v LIKE ?", key)] = []any{fmt.Sprintf("%v%%", value)}
	return s
}

func (s Query) EndWith(key string, value string) Query {
	s[fmt.Sprintf("%v LIKE ?", key)] = []any{fmt.Sprintf("%%%v", value)}
	return s
}

func (s Query) Cmp(key string, op string, value any) Query {
	s[fmt.Sprintf("%v %v ?", key, op)] = []any{value}
	return s
}

func (s Query) In(key string, values ...any) Query {
	s[fmt.Sprintf("%v IN ?", key)] = []any{values}
	return s
}

func (s Query) Custom(key string, values ...any) Query {
	s[key] = values
	return s
}
