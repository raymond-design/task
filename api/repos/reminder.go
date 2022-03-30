package repos

import (
	"encoding/json"
	"io"

	"github.com/raymond-design/task/api"
	"github.com/raymond-design/task/api/models"
	"github.com/raymond-design/task/api/services"
)

// the JSON db type
type FileDB interface {
	io.ReadWriter
	api.Stopper
	Size() int
	GenerateID() int
}

// each reminders repo
type Reminders struct {
	DB FileDB
}

// creates a new instance of Reminder repository
func NewReminders(db FileDB) *Reminders {
	return &Reminders{
		DB: db,
	}
}

// take sanpshot of current DB state
func (r Reminders) Save(reminders []models.Reminder) (int, error) {
	bs, err := json.Marshal(reminders)
	if err != nil {
		return 0, err
	}
	n, err := r.DB.Write(bs)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// filters reminders by a filtering function
func (r Reminders) Filter(filterFn func(reminder models.Reminder) bool) (services.RemindersMap, error) {
	bs := make([]byte, r.DB.Size())
	n, err := r.DB.Read(bs)
	if err != nil {
		e := models.WrapError("could not read from db", err)
		return services.RemindersMap{}, e
	}

	var reminders []models.Reminder
	err = json.Unmarshal(bs[:n], &reminders)
	if err != nil {
		e := models.WrapError("could not unmarshal json", err)
		return services.RemindersMap{}, e
	}

	res := services.RemindersMap{}
	for i, reminder := range reminders {
		if filterFn == nil || filterFn(reminder) {
			reminderMap := map[int]models.Reminder{}
			reminderMap[i] = reminder
			res[reminder.ID] = reminderMap
		}
	}
	return res, nil
}

// fetches the next DB AUTOINCREMENT id
func (r Reminders) NextID() int {
	return r.DB.GenerateID()
}
