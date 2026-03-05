package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

// APIHandler is the struct defining the handlers for the REST API
type APIHandler struct {}

// HostsHandler handles requests for the /hosts endpoint
func (h *APIHandler) HostsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the /hosts endpoint"))
}

// AppsHandler handles requests for the /apps endpoint
func (h *APIHandler) AppsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the /apps endpoint"))
}

// CISResultsHandler handles requests for the /cis-results endpoint
func (h *APIHandler) CISResultsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the /cis-results endpoint"))
}

// RegisterRoutes registers the API routes
func (h *APIHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hosts", h.HostsHandler).Methods(http.MethodGet)
	r.HandleFunc("/apps", h.AppsHandler).Methods(http.MethodGet)
	r.HandleFunc("/cis-results", h.CISResultsHandler).Methods(http.MethodGet)
}