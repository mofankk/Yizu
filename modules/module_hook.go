package modules

import (
	"gorm.io/gorm"
	"time"
)

func (h *House) BeforeCreate(tx *gorm.DB) (err error) {
	nt := time.Now().Format("2006-01-02 15:04:05")
	h.CreateTime = nt
	h.UpdateTime = nt
	return
}
