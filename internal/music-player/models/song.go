package models

import "time"

type Song struct {
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Duration time.Duration `json:"duration"`
}
