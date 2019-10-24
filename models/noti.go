package models

import "encoding/json"

// NotificationPayload notifications sent bu suunto
type NotificationPayload struct {
	UserName  string `form:"username" db:"_"`
	WorkoutID string `form:"workoutid" db:"_"`
}

func (payload NotificationPayload) String() string {
	ju, _ := json.Marshal(payload)
	return string(ju)
}
