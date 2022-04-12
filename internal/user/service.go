package user

import (
	"context"
	"rest-api-tutorial/pgk/logging"
)

type Service struct {
	Storage Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDto) (u User, err error) {
	// TODO for next one
	return
}
