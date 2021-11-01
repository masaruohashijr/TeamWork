package main

import (
	"fmt"
	"golang-interview-project-masaru-ohashi/cmd/common"
	"golang-interview-project-masaru-ohashi/pkg/api/REST"
	rmo "golang-interview-project-masaru-ohashi/pkg/repository/mongo"
	"golang-interview-project-masaru-ohashi/pkg/team"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

func main() {

	repo, err := rmo.NewMongoRepository(common.MONGO_URL, common.MONGO_DB, common.MONGO_TIMEOUT)
	if err != nil {
		log.Fatal(err)
	}
	service := team.NewMemberService(repo)
	handler := REST.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/Members", handler.GetAll)
	r.Get("/Member/{name}", handler.Get)
	r.Post("/Member", handler.Post)
	r.Put("/Member", handler.Put)
	r.Delete("/Member", handler.Delete)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :" + common.REST_PORT)
		errs <- http.ListenAndServe(":"+common.REST_PORT, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()
	fmt.Printf("Terminated %s", <-errs)
}
