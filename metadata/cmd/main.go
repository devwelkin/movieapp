package cmd

import (
	"log"
	"net/http"

	"github.com/devwelkin/movieapp/metadata/internal/controller/metadata"
	httphandler "github.com/devwelkin/movieapp/metadata/internal/handler/http"
	"github.com/devwelkin/movieapp/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting movie metadata sevice...")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
