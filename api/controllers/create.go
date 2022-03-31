package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/raymond-design/task/api/json_transport"
	"github.com/raymond-design/task/api/models"
	"github.com/raymond-design/task/api/services"
)

type creator interface {
	Create(reminderBody services.ReminderCreateBody) (models.Reminder, error)
}

func createReminder(service creator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Title    string        `json:"title"`
			Message  string        `json:"message"`
			Duration time.Duration `json:"duration"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			json_transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
			return
		}
		reminder, err := service.Create(services.ReminderCreateBody{
			Title:    body.Title,
			Message:  body.Message,
			Duration: body.Duration,
		})
		if err != nil {
			json_transport.SendError(w, err)
			return
		}
		json_transport.SendJSON(w, reminder, http.StatusCreated)
	})
}
