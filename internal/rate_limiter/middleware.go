package ratelimiter

import (
	"fmt"
	"net/http"
	"time"
)

const (
	LIMIT = 5
)

func (s *service) limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.Header.Get("X-Forwarded-For")
		ipParent, err := s.ipParser.GetParentIp(ipAddress)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println(ipAddress)
		ok := s.getVisitor(ipParent)
		if !ok {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *service) getVisitor(ip string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.visitors[ip]
	if !ok {
		s.visitors[ip] = NewVisitor()
		return true
	}
	if v.limit == LIMIT {
		if !v.cooldown {
			v.cooldown = true
			v.lastSeen = time.Now()
		}
		return false
	}
	v.limit++
	return true
}

func (s *service) cleanupVisitors() {
	for {
		s.mu.Lock()
		for _, v := range s.visitors {
			if time.Since(v.lastSeen) > 10*time.Second {
				v.lastSeen = time.Now()
				v.limit = 0
				v.cooldown = false
			}
		}
		s.mu.Unlock()
	}
}
