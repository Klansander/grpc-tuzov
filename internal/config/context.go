package config

import (
	"context"
	"fmt"

	parent "sso/pkg/context"

	"github.com/sirupsen/logrus"
)

type ctxConfig struct{}

// ContextWithConfig adds config to context
func ContextWithConfig(ctx context.Context, c *Config) context.Context {

	fmt.Println("!!!!!! ContextWithConfig")
	return context.WithValue(ctx, ctxConfig{}, c)

}

// GetConfig configFromContext returns config from context
func GetConfig(ctx context.Context) *Config {

	if c, ok := ctx.Value(ctxConfig{}).(*Config); ok {
		return c
	}

	if p, ok := ctx.Value(parent.ParentContext).(context.Context); ok {
		if c, ok := p.Value(ctxConfig{}).(*Config); ok {
			return c
		}
	}

	logrus.Fatal("Отсутствует инициализация системы конфигурации")
	return nil

}
