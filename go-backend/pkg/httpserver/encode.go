package httpserver

import (
	"encoding/json"
	"net/http"

	"go-backend/pkg/log"
)

func RespondJSON(status int, w http.ResponseWriter, r *http.Request, resp interface{}) {
	ctx := r.Context()

	logger := log.GetLogger(ctx)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		logger.WithError(err).Error("error writing json response")
	}
}
