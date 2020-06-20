package api

import (
	"net/http"
	"time"

	"go-backend/pkg/httpserver"
	"go-backend/pkg/log"
)

type HeathCheckResponse struct {
	CurrentTime time.Time
	Message     string
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.GetLogger(ctx)

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		logger.WithError(err).Error("error loading location")
	}

	resp := &HeathCheckResponse{
		CurrentTime: time.Now().In(loc),
		Message:     "hello there!",
	}

	httpserver.RespondJSON(http.StatusOK, w, r, resp)
}
