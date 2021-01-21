package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	router := chi.NewRouter()

	ir := &IndexRouter{}
	ar := &AboutRouter{}

	router.Mount("/", ir.Routes())
	router.Mount("/about", ar.Routes())

	return router
}
