package controllers

import (
	"net/http"

	"github.com/raymond-design/task/api/json_transport"
	"github.com/raymond-design/task/api/models"
)

type fetcher interface {
	Fetch(ids []int) ([]models.Reminder, error)
}

func fetchReminders(service fetcher) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ids, err := parseIDsParam(r.Context())
		if err != nil {
			json_transport.SendError(w, err)
			return
		}
		reminders, err := service.Fetch(ids)
		if err != nil {
			json_transport.SendError(w, err)
			return
		}
		json_transport.SendJSON(w, reminders, http.StatusOK)
	})
}
