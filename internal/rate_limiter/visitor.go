package ratelimiter

import (
	"time"
)

type visitor struct {
	lastSeen time.Time
	limit    int
}

func NewVisitor() *visitor {
	return &visitor{
		lastSeen: time.Now(),
		limit:    1,
	}
}

func freeVisitor(in *visitor) {
	in.lastSeen = time.Now()
	in.limit = 0
}
