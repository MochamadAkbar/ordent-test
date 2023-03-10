package main

import (
	"context"
	"fmt"
	"github.com/MochamadAkbar/ordent-test/handler"
	"log"
	"os"

	"github.com/MochamadAkbar/ordent-test/common/colorize"
	"github.com/MochamadAkbar/ordent-test/common/constants"
	"github.com/MochamadAkbar/ordent-test/config"
	"github.com/MochamadAkbar/ordent-test/injector"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func main() {
	ctx := context.Background()

	urlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db := config.NewDB(ctx, urlString)
	defer db.Close()

	router := config.NewRouter()
	router.Route("/api", func(r chi.Router) {
		_ = injector.InitializeUserService(db, r)
		r.Mount("/", handler.NewCartHandler())
	})

	server := config.NewServer(router)

	badge := colorize.MessageColorized(constants.Green, "ready")
	msg := fmt.Sprintf("[%s] started serve on [::]%s", badge, ":5000")
	log.Println(msg)
	if err := server.ListenAndServe(); err != nil {
		badge = colorize.MessageColorized(constants.Red, "stop")
		msg = fmt.Sprintf("[%s] server failed to start : %s", badge, err.Error())
		log.Fatalln(msg)
	}
}
