package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lucabecci/docker-golang/pkg"
)

type IndexRouter struct{}

func (ir *IndexRouter) Say(w http.ResponseWriter, r *http.Request) {

	information := pkg.Map{
		"message": "Hello my friend, welcome to my API",
		"aboutMe": "http://localhost:3000/about",
	}
	pkg.JSON(w, r, http.StatusOK, information)
}

func (ir *IndexRouter) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", ir.Say)
	return router
}
