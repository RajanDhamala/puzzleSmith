package Routes

import (
	"encoding/json"
	"net/http"

	"github.com/RajanDhamala/puzzleSmith/Controllers"
	"github.com/go-chi/chi/v5"
)

func GameRouter(r chi.Router, controller *Controllers.Controller) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		repsonse := map[string]string{
			"message": "gamer route is up and running",
		}

		json.NewEncoder(w).Encode(repsonse)
	})

	r.Get("/process", controller.ProcessGame)
}
