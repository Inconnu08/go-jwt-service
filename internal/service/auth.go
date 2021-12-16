package service

import (
	"fmt"
	"io/ioutil"
	"sync/atomic"
	"time"
)


func (s *Service) CreateToken(subject string) (string, string, error) {
	// for stats
	s.addStatValue(subject + "-auth", 1)

	// Create a new JWT token, should expire exactly 24 hours
	tok, err := s.token.Create(time.Hour*24, subject)
	if err != nil {
		return "", "", err
	}

	return tok, string(s.token.PublicKey), nil
}

func (s *Service) VerifyToken(token string) (string, error) {
	content, err := s.token.Validate(token)
	if err != nil {
		return "", err
	}

	username, ok := content.(string)
	if !ok {
		return "", err
	}

	// for stats
	s.addStatValue(username + "-verify", 1)

	return username, nil
}

func (s *Service) GetReadme() (string, error) {
	readme, err := ioutil.ReadFile("./README.txt")
	if err != nil {
		return "", err
	}

	return string(readme), nil
}

func (s *Service) getStatValue(key string) (int64, bool) {
	count, ok := s.statsDB.Load(key)
	if ok {
		return atomic.LoadInt64(count.(*int64)), true
	}
	return 0, false
}

func (s *Service) addStatValue(key string, value int64) int64 {
	count, loaded := s.statsDB.LoadOrStore(key, &value)
	if loaded {
		return atomic.AddInt64(count.(*int64), value)
	}
	return *count.(*int64)
}

// GetStats tells the authenticated user how many times they have visited /auth and /verify explicitly and implicitly
func (s *Service) GetStats(username string) (string, error) {
	visitedAuth, _ := s.getStatValue(username + "-auth")
	visitedVerify, _ := s.getStatValue(username + "-verify")

	body := fmt.Sprintf("visited /auth: %d times(s)\n visited /verify: %d times(s)\n", visitedAuth, visitedVerify)

	return body, nil
}
