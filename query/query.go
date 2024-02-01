package query

import (
	"fmt"
	"github.com/qwerty22121998/gobase/tlf"
	"gorm.io/gorm"
)

type Condition func(db *gorm.DB) *gorm.DB

func group(c ...Condition) []func(db *gorm.DB) *gorm.DB {
	return tlf.Map(c, func(condition Condition) func(db *gorm.DB) *gorm.DB {
		return condition
	})
}

func Or(c ...Condition) Condition {
	return func(db *gorm.DB) *gorm.DB {
		tx := db.Session(&gorm.Session{NewDB: true}).Unscoped()
		for _, condition := range c {
			tx2 := tx.Session(&gorm.Session{NewDB: true}).Unscoped()
			tx.Or(tx2.Scopes(condition))
		}
		return db.Where(tx)
	}
}

func And(c ...Condition) Condition {
	return func(db *gorm.DB) *gorm.DB {
		tx := db.Session(&gorm.Session{NewDB: true}).Unscoped()
		return db.Where(
			tx.Scopes(group(c...)...),
		)
	}
}

func NonZero(value any, condition Condition) Condition {
	return func(db *gorm.DB) *gorm.DB {
		if isZero(value) {
			return db
		}
		return condition(db)
	}
}

func Equal(key string, value any) Condition {
	return Cmp(key, "=", value)
}

func Like(key string, value string) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v LIKE ?", key), fmt.Sprintf("%%%v%%", value))
	}
}

func StartWith(key string, value string) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v LIKE ?", key), fmt.Sprintf("%v%%", value))
	}
}

func EndWith(key string, value string) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v LIKE ?", key), fmt.Sprintf("%%%v", value))
	}
}

func Cmp(key string, op string, value any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v %v ?", key, op), value)
	}
}
func In(key string, values ...any) Condition {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%v IN (?)", key), values)
	}
}
