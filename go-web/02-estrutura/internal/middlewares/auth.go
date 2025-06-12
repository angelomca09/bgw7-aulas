package middlewares

import (
	"encoding/json"
	"estrutura-api/config"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		apiToken := config.GetAPIConfig().ApiToken

		if token != apiToken {
			code := http.StatusUnauthorized // 401
			body := struct {
				Message string `json:"message"`
			}{
				"Unauthorized",
			}

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(body)
			return
		}

		next.ServeHTTP(w, r)
	})
}
