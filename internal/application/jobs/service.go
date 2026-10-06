package jobs

import "log/slog"

// Service - Jobs use cases, consumed by the HTTP handlers
type Service struct {
	repo JobRepository
	log  *slog.Logger
}

func NewService(repo JobRepository, log *slog.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
