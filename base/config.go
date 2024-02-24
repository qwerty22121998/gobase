package base

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/samber/do/v2"
)

func ParseConfigInject[T any](do.Injector) (T, error) {
	return ParseConfig[T]()
}

func ParseConfig[T any]() (T, error) {
	return ParseConfigWithPrefix[T]("")
}

func ParseConfigWithPrefix[T any](prefix string) (T, error) {
	var cfg T
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
