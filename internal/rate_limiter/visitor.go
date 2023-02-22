package ratelimiter

import (
	"time"
)

type visitor struct {
	lastSeen time.Time
	limit    int
	cooldown bool
}

func NewVisitor() *visitor {
	return &visitor{
		lastSeen: time.Now(),
		limit:    1,
		cooldown: false,
	}
}
