package preload

import "gorm.io/gorm"

type PreloadType string

const (
	NORMAL PreloadType = "PRELOAD"
	JOIN   PreloadType = "JOIN"
)

type Opt interface {
	Apply(db *gorm.DB) *gorm.DB
}

type OptGroup []Opt

func (g OptGroup) Apply(db *gorm.DB) *gorm.DB {
	for _, opt := range g {
		db = opt.Apply(db)
	}
	return db
}

type preload struct {
	field string
	ptype PreloadType
	args  []any
}

func Group(opts ...Opt) Opt {
	return OptGroup(opts)
}

func (p *preload) Apply(db *gorm.DB) *gorm.DB {
	switch p.ptype {
	case NORMAL:
		return db.Preload(p.field, p.args...)
	case JOIN:
		return db.Joins(p.field, p.args...)
	default:
		return db
	}
}

func Preload(field string, args ...any) Opt {
	return &preload{
		field: field,
		ptype: NORMAL,
		args:  args,
	}
}

func Join(field string, args ...any) Opt {
	return &preload{
		field: field,
		ptype: JOIN,
		args:  args,
	}
}
