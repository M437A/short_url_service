package routers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"short_url/iternal/controllers"
)

func CreateRouters(router *chi.Mux) {
	router.Route("/", func(r chi.Router) {
		r.Post("/", controllers.CreateShortUrl)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	controllers.Router = router
}
