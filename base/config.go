package base

import "github.com/kelseyhightower/envconfig"

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
