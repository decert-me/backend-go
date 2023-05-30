package dao

import (
	"backend-go/internal/app/model"
	"time"
)

func (d *Dao) GetUserResourceBeforeList(before time.Time) {
	d.db.Model(&model.Users{}).Where("resource_time < ?", before)
}
