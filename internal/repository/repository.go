package repository

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type UserRepository interface {
	GetVal(key string) (string, error)
	SetVal(key, val string, ttl time.Duration) (string, error)
	AddVal(key, val string) (int64, error)
	DelVal(key string) (int64, error)
	KeysVal(key string) ([]string, error)
}

type repository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) (*repository, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}

	return &repository{
		db: db,
	}, nil
}

func (r *repository) GetVal(key string) (string, error) {
	return r.db.Get(key).Result()
}

func (r *repository) SetVal(key, val string, ttl time.Duration) (string, error) {

	return r.db.Set(key, val, ttl*time.Second).Result()
}

func (r *repository) AddVal(key, val string) (int64, error) {
	return r.db.Append(key, val).Result()
}

func (r *repository) DelVal(key string) (int64, error) {
	return r.db.Del(key).Result()
}
func (r *repository) KeysVal(key string) ([]string, error) {
	return r.db.Keys(key).Result()
}
