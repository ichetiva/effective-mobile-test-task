package responder

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, data interface{})
	ErrorJSON(w http.ResponseWriter, message string, status int)
}

type Respond struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) Responder {
	return &Respond{
		log: log,
	}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		r.log.Errorf("failed output json: %v\n", err)
	}
}

func (r *Respond) ErrorJSON(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(ErrorResponse{Message: message}); err != nil {
		r.log.Errorf("failed output error json: %v\n", err)
	}
}
