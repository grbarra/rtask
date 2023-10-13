package models

import "time"

type User struct {
	Key  string        `json:"key"`
	Name string        `json:"name"`
	TTL  time.Duration `json:"ttl,omitempty"`
}

type UserKeys struct {
	Keys []string `json:"key"`
}
