package main

import (
	"fmt"
	"golang-interview-project-masaru-ohashi/pkg/api"
	"golang-interview-project-masaru-ohashi/pkg/member"
	rm "golang-interview-project-masaru-ohashi/pkg/repository/mongo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func main() {
	mongoURL := "mongodb://localhost:27017/"
	mongodb := "mongodb"
	mongoTimeout := 60000
	repo, err := rm.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}
	service := member.NewMemberService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{name}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(":8000", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}
