package server

import "github.com/ehsundar/deeds/internal/storage"

type Server struct {
	Store storage.Storage
}
