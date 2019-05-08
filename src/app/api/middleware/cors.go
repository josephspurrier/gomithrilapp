package middleware

import (
	"net/http"
)

// CORS will handle pre-flight OPTIONS requests.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// These were the old settings.
		/*router.SetGlobalCors(&vestigo.CorsAccessControl{
			AllowOrigin:  []string{"*"},
			AllowHeaders: []string{"Content-Type", "Origin", "X-Requested-With", "Accept"},
		})*/

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
