package main

import (
	"log"

	"github.com/vladjong/rate_limiter/internal/config"
	ipparser "github.com/vladjong/rate_limiter/internal/ip_parser"
	ratelimiter "github.com/vladjong/rate_limiter/internal/rate_limiter"
)

func main() {
	// пакет для работы с переменными окружения
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	// пакет для работы с ip адрессами
	ipParser := ipparser.New(cfg.Prefix)

	// основной сервис
	rate := ratelimiter.New(ipParser, *cfg)
	rate.Run()
}
