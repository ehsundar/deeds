package server

import (
	"github.com/ehsundar/deeds/internal/server/templates"
	"net/http"
)

func (s *Server) HandleView(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token not provided", http.StatusBadRequest)
		return
	}

	err := templates.ViewTemplateParams{ImageFileName: token}.Render(w)
	if err != nil {
		http.Error(w, "could not render template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
