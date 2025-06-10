Ce dossier contient les gestionnaires HTTP qui traitent les requêtes entrantes et renvoient les réponses appropriées

.package handlers

import (
	"net/http"
	"mon-projet/internal/services"
)

type ExampleHandler struct {
	exampleService *services.ExampleService
}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{
		exampleService: services.NewExampleService(),
	}
}

func (h *ExampleHandler) HandleExample(w http.ResponseWriter, r *http.Request) {
	// Logique pour traiter la requête
	example, err := h.exampleService.GetExample()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renvoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(example)
}