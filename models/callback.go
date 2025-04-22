package models

import (
	"time"
)

type Callback struct {
	ID        int64     `json:"id"`
	IP        string    `json:"ip"`
	Domain    string    `json:"domain"`
	Params    string    `json:"params"`
	CreatedAt time.Time `json:"created_at"`
}

type CallbackResponse struct {
	Total     int64      `json:"total"`
	Callbacks []Callback `json:"callbacks"`
}
