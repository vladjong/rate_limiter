package main

import (
	"log"

	"github.com/vladjong/rate_limiter/internal/config"
	ipparser "github.com/vladjong/rate_limiter/internal/ip_parser"
	ratelimiter "github.com/vladjong/rate_limiter/internal/rate_limiter"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ipParser := ipparser.New(24)

	rate := ratelimiter.New(ipParser, *cfg)
	rate.Run()
}
