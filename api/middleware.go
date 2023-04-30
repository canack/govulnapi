package api

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func (s *Api) userDispatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			_, creds, _ = jwtauth.FromContext(r.Context())
			user_id     = int(creds["user_id"].(float64))
		)

		user, err := s.db.GetUserById(user_id)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
