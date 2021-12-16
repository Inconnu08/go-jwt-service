package service

import (
	"sync"

	"socialmedia-auth/internal/service/token"
)

// contains the core business logic.
type Service struct {
	token 		token.JWT
	statsDB 	sync.Map // this map stores all information related to stats
}

func New(token token.JWT) *Service {
	return &Service{token: token}
}
