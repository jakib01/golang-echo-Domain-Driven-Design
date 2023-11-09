package Common

import "time"

type Times struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time"`
}
