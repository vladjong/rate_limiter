package ratelimiter

import (
	"net/http"
	"time"
)

// middleware для ограничения количества запросов
func (s *service) limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// определяем ip по заголовку X-Forwarded-For
		ipAddress := r.Header.Get("X-Forwarded-For")
		// определяем подсеть, полученного ip
		ipParent, err := s.ipParser.GetParentIp(ipAddress)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		// проверка на кол-ва запросов к сервису
		ok := s.getVisitor(ipParent)
		if !ok {
			// отправляем ответ по RFC 6585 (429 To Many Requests etc.)
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// метод для подсчета запросов к сервису
func (s *service) getVisitor(ip string) bool {
	// лок для параллейного доступа к мапе
	s.mu.Lock()
	defer s.mu.Unlock()
	// проверяем есть ли у нас в мапе эл-ет с подсетью
	v, ok := s.visitors[ip]
	if !ok {
		// создаем нового пользователя
		s.visitors[ip] = NewVisitor()
		return true
	}
	// проверка на max кол-во запросов к сервису
	if v.limit == s.cfg.Limit {
		// если не было cooldown, то обнуляем время
		if !v.cooldown {
			v.lastSeen = time.Now()
			v.cooldown = true
		}
		return false
	}
	// добавляем к счетчику запросов +1
	v.limit++
	return true
}

// метод для обновления доступа к сервису (работает в фоновом режиме)
func (s *service) cleanupVisitors() {
	for {
		// лок для параллейного доступа к мапе
		s.mu.Lock()
		for _, v := range s.visitors {
			// проверка на последний запрос к сервису
			if time.Since(v.lastSeen) > time.Duration(s.cfg.TimeCooldown)*time.Minute {
				freeVisitor(v)
			}
		}
		s.mu.Unlock()
	}
}
