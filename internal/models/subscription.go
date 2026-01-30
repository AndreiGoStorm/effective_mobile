package models

import "time"

type Subscription struct {
	ID          int64      `fake:"{number:1,100000}"`
	ServiceName string     `json:"service_name" fake:"{appname}"`
	Price       int        `fake:"{number:100,5000}"`
	UserUUID    string     `fake:"{uuid}"`
	StartDate   time.Time  `fake:"{skip}"`
	EndDate     *time.Time `fake:"{skip}"`
	CreatedAt   time.Time  `fake:"{skip}"`
	UpdatedAt   time.Time  `fake:"{skip}"`
}
