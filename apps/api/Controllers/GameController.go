package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RajanDhamala/puzzleSmith/ProcessPipline"
	"github.com/RajanDhamala/puzzleSmith/Utils"
)

func (ctrl *Controller) ProcessGame(w http.ResponseWriter, r *http.Request) {
	name := "NbcWala"
	games, err := utils.FetchProcess(name)
	if err != nil {
		fmt.Println("error while fetching game", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to find games",
		})
		return
	}
	processpipline.StartProcessingGame(name, games, ctrl.stockfish)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "game being processed in bg",
	})
}
