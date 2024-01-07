package zap

import (
	"go.uber.org/zap"

	"github.com/imthaghost/musik/backend/internal/logger"
)

type Service struct {
	Log *zap.Logger
}

// Msg ...
func (s *Service) Msg(message string) {
	s.Log.Info(message)
}

// Error ...
func (s *Service) Error(message string) {
	s.Log.Error(message)
}

// NewService creates a
func NewService() logger.Service {
	log, _ := zap.NewProduction()
	defer log.Sync()

	return &Service{
		Log: log,
	}
}
