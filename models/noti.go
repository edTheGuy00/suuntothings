package models

import "encoding/json"

// NotificationPayload notifications sent bu suunto
type NotificationPayload struct {
	UserName  string `json:"username" db:"_"`
	WorkoutID string `json:"workoutid" db:"_"`
}

func (u NotificationPayload) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}
