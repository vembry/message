package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("starting api...")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		raw, _ := json.Marshal(map[string]interface{}{
			"status":  true,
			"message": "ok",
		})

		// write to response
		w.Header().Set("Content-Type", "application/json")
		w.Write(raw)
	})

	httpserver := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	if err := httpserver.ListenAndServe(); err != nil {
		log.Fatalf("error on serving http-server. err=%v", err)
	}

	slog.Info("shutting down api...")
}
