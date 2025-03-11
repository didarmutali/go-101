package main

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"main.go/internal/http"
	"main.go/internal/store/postgres"
)

func main() {
	urlExample := "postgres://localhost:5432/go-101"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	cache, err := lru.New2Q(6)
	if err != nil {
		panic(err)
	}

	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8080"),
		http.WithStore(store),
		http.WithCache(cache),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
