package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Gurtinho/virtualroom/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type API struct {
	q *db.Queries
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func main() {
	_ = godotenv.Load()

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	_ = &API{q: db.New(pool)}

	mux := http.NewServeMux()

	// cria os handlers

	log.Println("rodando em :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
