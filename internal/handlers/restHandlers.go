package handlers

import (
	"encoding/json"
	"net/http"

	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/helpers"
	"gerrit.ericsson.se/a/DETES/com.ericsson.de.stsoss/inventory-app/internal/models"
	"github.com/go-chi/chi/v5"
)

func (m *Repository) GetAllDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allDeployments := helpers.GetAllFromCollectiion("Deployments", m.App)
	json.NewEncoder(w).Encode(allDeployments)
}

func (m *Repository) GetOneDeployment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "GET")
	id := chi.URLParam(r, "id")
	deployment, err := helpers.GetOneDeployment(id, m.App)
	if err != nil {
		json.NewEncoder(w).Encode("no result")
		return
	}
	json.NewEncoder(w).Encode(deployment)
}

func (m *Repository) CreateDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "POST")

	var deployment models.Deployment
	_ = json.NewDecoder(r.Body).Decode(&deployment)
	helpers.InsertOneDeployment(deployment, m.App)
	json.NewEncoder(w).Encode(deployment)
}

func (m *Repository) MarkDeploymentAsUsed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "PUT")
	id := chi.URLParam(r, "id")
	helpers.UpdateOneDeployment(id, m.App)
	json.NewEncoder(w).Encode(id)

}

func (m *Repository) DeleteDeployment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Controll-Allow-Methods", "DELETE")
	id := chi.URLParam(r, "id")
	helpers.DeletOneDeployment(id, m.App)
	json.NewEncoder(w).Encode(id)

}
