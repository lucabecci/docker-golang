package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lucabecci/docker-golang/pkg"
)

type AboutRouter struct{}

func (ar *AboutRouter) Creator(w http.ResponseWriter, r *http.Request) {

	information := pkg.Map{
		"name":     "Luca Becci",
		"github":   "https://github.com/lucabecci",
		"linkedin": "https://www.linkedin.com/in/luca-becci-b8044b198/",
	}
	pkg.JSON(w, r, http.StatusOK, information)
}

func (ar *AboutRouter) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", ar.Creator)
	return router
}
