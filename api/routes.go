package api

import (
	_ "govulnapi/api/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Api) setupRoutes() {
	r := s.router

	// CWE-942: Permissive Cross-domain Policy with Untrusted Domains
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Mount("/", httpSwagger.WrapHandler)

	r.Route("/api", func(r chi.Router) {
		r.Get("/coins", s.getCoins)

		// CWE-598: Use of GET Request Method With Sensitive Query Strings
		r.Get("/register", s.registerUser)
		r.Get("/login", s.loginUser)

		// Token needed
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(s.jwtAuth))
			r.Use(jwtauth.Authenticator)
			r.Use(s.userDispatcher)

			r.Get("/balances/coin", s.getCoinBalances)
			r.Get("/balances/usd", s.getUsdBalances)

			r.Post("/orders", s.addOrder)
			r.Get("/orders", s.getOrders)

			r.Get("/transactions", s.getTransactions)
			r.Post("/transactions", s.addTransaction)

			r.Put("/user/email", s.updateEmail)
			r.Put("/user/password", s.updatePassword)
		})
	})

}
