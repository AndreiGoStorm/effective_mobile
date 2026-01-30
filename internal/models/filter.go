package models

import "time"

type Filter struct {
	ServiceName string
	UserUUID    string
	StartDate   time.Time
	EndDate     *time.Time
}
