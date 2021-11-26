package application

import (
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/ports"
)

type service struct {
	r ports.Repository
	i ports.ImageRepository
}

func NewService(repository ports.Repository, imageRepository ports.ImageRepository) *service {
	return &service{
		r: repository,
		i: imageRepository,
	}
}

func (s *service) Ping() {
	s.r.Ping()
}
