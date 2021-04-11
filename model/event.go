package model

import "time"

type Event struct {
	Metadata  []byte    `json:"metadata"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
