package context

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ParentContext = "parentContext"

// ContextWithParentContext adds config to context
func ContextWithParentContext(ctx *gin.Context, ctxParent context.Context) {

	ctx.Set(ParentContext, ctxParent)

}

// GetParentContext returns config from context
func GetParentContext(ctx context.Context) context.Context {

	if c, ok := ctx.Value(ParentContext).(context.Context); ok {
		return c
	}

	logrus.Error("Отсутствует родительский контекст")

	return context.Background()

}
