package server

import (
	"github.com/ehsundar/deeds/internal/server/templates"
	"io"
	"net/http"
	"os"
	"path"
)

func (s *Server) HandleFrom(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "token not provided", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		err := templates.FormTemplateParams{}.Render(w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	if r.Method == http.MethodPost {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "token not provided", http.StatusBadRequest)
			return
		}

		file, _, err := r.FormFile("img")
		if err != nil {
			http.Error(w, "Failed to retrieve the uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a new file on the server to Store the uploaded file
		uploadedFile, err := os.Create(path.Join("images", token))
		if err != nil {
			http.Error(w, "Failed to create the file on the server", http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()

		// Copy the uploaded file's content to the newly created file
		_, err = io.Copy(uploadedFile, file)
		if err != nil {
			http.Error(w, "Failed to copy the file content", http.StatusInternalServerError)
			return
		}

		s.attachAddon(r.Context(), token)

		http.Redirect(w, r, "/confirm", http.StatusFound)
	}
}
