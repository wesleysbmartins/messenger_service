package exceptions

import (
	"context"
	"fmt"
	"messenger_service/internal/shared/logger"
)

type DatabaseException struct{}

type IDatabaseException interface {
	Handle(ctx context.Context, operation string, filter any, value any, err error)
}

func (e *DatabaseException) Handle(ctx context.Context, operation string, filter any, value any, err error) {
	log := &logger.Logger{}

	context := fmt.Sprintf("%s", ctx.Value("Database"))
	message := err.Error()

	toLog := fmt.Sprintf("Context: %s\nMessage: %s\nOperation: %s\nFilter: %s\nValue: %s\nError: %s\n", context, message, operation, filter, value, err)

	log.Write("DATABASE EXCEPTION", toLog)
}
