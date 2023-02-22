package main

import (
	ipparser "github.com/vladjong/rate_limiter/internal/ip_parser"
	ratelimiter "github.com/vladjong/rate_limiter/internal/rate_limiter"
)

func main() {
	ipParser := ipparser.New(24)
	rate := ratelimiter.New(ipParser)
	rate.Run()
}
