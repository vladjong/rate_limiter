package ratelimiter

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vladjong/rate_limiter/internal/config"
	ipparser "github.com/vladjong/rate_limiter/internal/ip_parser"
)

type service struct {
	ipParser ipparser.IpParser
	visitors (map[string]*visitor)
	mu       sync.Mutex
	cfg      config.Config
}

func New(ipParser ipparser.IpParser, cfg config.Config) *service {
	return &service{
		ipParser: ipParser,
		visitors: make(map[string]*visitor),
		cfg:      cfg,
	}
}

// метод запуск сервиса
func (s *service) Run() {
	// фоновая задача, которая обновляет доступ пользователя к сервису
	go s.cleanupVisitors()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// middleware для ограничения количества запросов
	r.Use(s.limit)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PING-PONG"))
	})

	r.Get("/refresh", s.Refresh)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
