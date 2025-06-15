package main

import (
	"github.com/WilliamKSilva/go-hexagonal/internal/adapters/http"
)

func main() {
	// Initialize HTTP server
	http.NewServer().Run()
}
