package controllers

import (
	"net/http"

	"github.com/raymond-design/task/api/json_transport"
)

type deleter interface {
	Delete(ids []int) error
}

func deleteReminders(service deleter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ids, err := parseIDsParam(r.Context())
		if err != nil {
			json_transport.SendError(w, err)
			return
		}
		err = service.Delete(ids)
		if err != nil {
			json_transport.SendError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
