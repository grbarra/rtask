package service

import (
	"fmt"
	"log"
	"time"

	"gitlab.com/grbarra/rtask/internal/models"
	"gitlab.com/grbarra/rtask/internal/repository"
)

type UserService interface {
	GetUser(key string) (string, error)
	SetUser(key, val string, ttl time.Duration) (string, error)
	AddUser(key, value string) (int64, error)
	DelUser(key string) (int64, error)
	KeysUser(key string) (models.UserKeys, error)
}

type service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) *service {
	return &service{
		repo: repo,
	}
}
func (s *service) GetUser(key string) (string, error) {
	val, err := s.repo.GetVal(key)
	if err != nil {
		return "", fmt.Errorf("error inserting URL: %s", err)
	}

	return val, nil
}

func (s *service) SetUser(key, value string, ttl time.Duration) (string, error) {
	val, err := s.repo.SetVal(key, value, ttl)
	if err != nil {
		log.Fatal(err)
	}
	return val, nil
}

func (s *service) AddUser(key, value string) (int64, error) {
	val, err := s.repo.AddVal(key, value)
	if err != nil {
		log.Fatal(err)
	}
	return val, nil
}

func (s *service) DelUser(key string) (int64, error) {
	val, err := s.repo.DelVal(key)
	if err != nil {
		log.Fatal(err)
	}
	return val, nil
}

func (s *service) KeysUser(key string) (models.UserKeys, error) {
	val, err := s.repo.KeysVal(key)
	if err != nil {
		log.Fatal(err)
	}
	resp := models.UserKeys{
		Keys: val,
	}
	return resp, nil
}
