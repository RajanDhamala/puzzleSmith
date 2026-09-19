package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/RajanDhamala/puzzleSmith/Controllers"
	"github.com/RajanDhamala/puzzleSmith/Database"
	opening "github.com/RajanDhamala/puzzleSmith/Opening"
	processpipline "github.com/RajanDhamala/puzzleSmith/ProcessPipline"
	"github.com/RajanDhamala/puzzleSmith/Routes"
	"github.com/RajanDhamala/puzzleSmith/internal/db"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dbPool, err := Database.ConnectDB()
	if err != nil {
		fmt.Println("error while connecting to database:", err)
		return
	}
	defer dbPool.Close()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		fmt.Println("PORT environment variable is not set")
		panic("PORT environment variable is not set")
	}

	go opening.ReadAllFiles()
	instance := processpipline.ConnectStockfish()

	controller := Controllers.NewController(db.New(dbPool), dbPool, instance)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "server is up my friend",
		})
	})

	r.Route("/user", func(r chi.Router) { Routes.UserRouter(r, controller) })
	r.Route("/game", func(r chi.Router) { Routes.GameRouter(r, controller) })

	fmt.Println("server is running on port", PORT)
	err = http.ListenAndServe(":"+PORT, r)
	if err != nil {
		fmt.Println("failed to start server", err.Error())
		panic("failed to start server")
	}
}
