package base_mapper

import (
	"github.com/qwerty22121998/gobase/base_dto"
	"github.com/qwerty22121998/gobase/base_model"
	"time"
)

type Base struct{}

func timeOrNil(ts int64) *time.Time {
	if ts == 0 {
		return nil
	}
	t := time.Unix(ts, 0)
	return &t
}

func (Base) ToModel(d base_dto.DTO) base_model.Model {
	return base_model.Model{
		ID:        d.ID,
		CreatedAt: time.Unix(0, d.CreatedAt),
		UpdatedAt: time.Unix(0, d.UpdatedAt),
		CreatedBy: d.CreatedBy,
		UpdatedBy: d.UpdatedBy,
	}
}

func (Base) ToDTO(m base_model.Model) base_dto.DTO {
	return base_dto.DTO{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
	}
}
