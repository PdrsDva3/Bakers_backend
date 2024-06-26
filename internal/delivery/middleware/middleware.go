package middleware

import "Bakers_backend/pkg/logger"

type Middleware struct {
	logger *logger.Logs
}

func InitMiddleware(logger *logger.Logs) Middleware {
	return Middleware{
		logger: logger,
	}
}
