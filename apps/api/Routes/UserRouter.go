package Routes

import (
	"github.com/RajanDhamala/puzzleSmith/Controllers"
	// "github.com/RajanDhamala/puzzleSmith/Middlewares"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router, ctrl *Controllers.Controller) {
	// r.Get("/me", middlewares.AuthMe)
	r.Post("/register", ctrl.RegisterUser)
	r.Post("/login", ctrl.LoginUser)
	r.Get("/logout", ctrl.LogoutUser)
}
