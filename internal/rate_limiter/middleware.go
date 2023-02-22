package ratelimiter

import (
	"net/http"
	"time"
)

func (s *service) limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.Header.Get("X-Forwarded-For")
		ipParent, err := s.ipParser.GetParentIp(ipAddress)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
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
	if v.limit == s.cfg.Limit {
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
			if v.cooldown && time.Since(v.lastSeen) > time.Duration(s.cfg.TimeCooldown)*time.Minute {
				v.lastSeen = time.Now()
				v.limit = 0
				v.cooldown = false
			} else if v.cooldown && time.Since(v.lastSeen) > time.Duration(s.cfg.TimeLimit)*time.Minute {
				v.lastSeen = time.Now()
				v.limit = 0
				v.cooldown = false
			}
		}
		s.mu.Unlock()
	}
}
