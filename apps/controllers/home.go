package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{Directory: "./templates",
		Layout:     "layout",
		Extensions: []string{".html"},
	})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "y",
		"body":  "y",
	})
}
