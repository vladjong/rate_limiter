package ratelimiter

import "net/http"

func (s *service) Refresh(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.Header.Get("X-Forwarded-For")
	// определяем подсеть, полученного ip
	ipParent, err := s.ipParser.GetParentIp(ipAddress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	s.mu.Lock()
	s.visitors[ipParent] = NewVisitor()
	s.mu.Unlock()
	w.Write([]byte("OK"))
}
