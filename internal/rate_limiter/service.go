package ratelimiter

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ipparser "github.com/vladjong/rate_limiter/internal/ip_parser"
)

type service struct {
	ipParser ipparser.IpParser
	visitors (map[string]*visitor)
	mu       sync.Mutex
}

func New(ipParser ipparser.IpParser) *service {
	return &service{
		ipParser: ipParser,
		visitors: make(map[string]*visitor),
	}
}

func (s *service) Run() {
	go s.cleanupVisitors()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(s.limit)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PING-PONG"))
	})
	http.ListenAndServe(":3000", r)
}
