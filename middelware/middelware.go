package middelware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("IP: %s | URL: %s | User Agent: %s | Method: %s", r.RemoteAddr, r.URL.Path, r.UserAgent(), r.Method)
	})
}
