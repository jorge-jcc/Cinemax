package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	imagerepository "github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/imageRepository"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/repository"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/adapters/server/util/token"
	"github.com/jorge-jcc/cinemax/cinemax-backend/internal/application"
)

func main() {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	repository := repository.NewRepository(db)
	imageRepository := imagerepository.NewImageRepository("/home")
	service := application.NewService(repository, imageRepository)
	token := token.NewJWTMaker(os.Getenv("secret"))
	server.NewHandler(router, service, token)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}

func getConnection() (*sqlx.DB, error) {
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"database", 5432, "cinemax", "cinemax", "cinemax",
	)
	return sqlx.Open("postgres", uri)
}
