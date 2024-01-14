package models

import (
	"time"

	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	UserId uint `json:"user_id"`
	Meta

	Anti          bool `json:"anti"`
	FrequencyType uint `json:"frequency_type"` // 1 - Daily, 2 - Weekly
	Target        uint `json:"target"`         // Limit in case of anti
	Mode          uint `json:"mode"`           // Times, Minutes, Label etc.
	Status        uint `json:"status"`         // 0 - Paused, 1 - Active
}

// will just have daily logs, can use group by for weekly goals
// or can cache or rewrite in table with different type
type HabitLog struct {
	gorm.Model

	HabitID     uint      `json:"habit_id"`
	Date        time.Time `json:"date"`
	ResultCount uint      `json:"result_count"`
}
