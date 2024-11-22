package query

import "gorm.io/gorm"

type Query interface {
	Apply(db *gorm.DB) *gorm.DB
}
