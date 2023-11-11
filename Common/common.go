package Common

import (
	"gorm.io/gorm"
	"time"
)

type Times struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time"`
}

type Repository interface {
	// TxStart
	TxStart() (*gorm.DB, error)
	// TxCommit
	TxCommit(tx *gorm.DB) error
	// TxRollback
	TxRollback(tx *gorm.DB)
}
