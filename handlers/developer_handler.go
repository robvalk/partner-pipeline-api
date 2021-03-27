package handlers

import (
	"net/http"

	"github.com/robvalk/partner-pipeline-api/models"
)

// DeveloperHandler implements ServerInterface
type DeveloperHandler struct {
}

func NewDeveloperHandler() *DeveloperHandler {
	return &DeveloperHandler{}
}

func (h *DeveloperHandler) SearchDevelopers(w http.ResponseWriter, r *http.Request, params models.SearchDevelopersParams)

func (h *DeveloperHandler) CreateDeveloper(w http.ResponseWriter, r *http.Request)

func (h *DeveloperHandler) DeleteDeveloper(w http.ResponseWriter, r *http.Request, developerId string)

func (h *DeveloperHandler) GetDeveloper(w http.ResponseWriter, r *http.Request, developerId string)

func (h *DeveloperHandler) UpdateDeveloper(w http.ResponseWriter, r *http.Request, developerId string)
