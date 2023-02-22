package ratelimiter

import (
	"time"
)

// структура для хранения:времени последнего изменения ,кол-ва запросов
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

func freeVisitor(in *visitor) {
	in.lastSeen = time.Now()
	in.limit = 0
	in.cooldown = false
}
