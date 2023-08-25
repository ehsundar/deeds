package server

import (
	"github.com/ehsundar/deeds/internal/server/templates"
	"net/http"
)

func (s *Server) HandleConfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		err := templates.ConfirmTemplateParams{}.Render(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}
}
