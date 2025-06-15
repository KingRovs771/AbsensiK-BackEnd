package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CORSMiddleware() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Full-Name"},
		AllowCredentials: true,
	}).Handler
}
