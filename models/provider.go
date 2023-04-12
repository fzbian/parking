package models

import "time"

type Provider struct {
	ID          int        `json:"id" db:"id"`
	PlateNumber string     `json:"plateNumber" db:"plateNumber"`
	EntryTime   time.Time  `json:"entryTime" db:"entryTime"`
	ExitTime    *time.Time `json:"exitTime" db:"exitTime"`
}
