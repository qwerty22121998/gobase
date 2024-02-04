package base_model

import (
	"gorm.io/gorm"
	"time"
)

type IModel interface {
	SetCreated(by string, at time.Time)
	SetUpdated(by string, at time.Time)
}

type Model struct {
	gorm.Model
	CreatedBy string
	UpdatedBy string
}

func (m *Model) SetCreated(by string, at time.Time) {
	m.CreatedBy = by
	m.CreatedAt = at
}

func (m *Model) SetUpdated(by string, at time.Time) {
	m.UpdatedBy = by
	m.UpdatedAt = at
}
