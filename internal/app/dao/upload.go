package dao

import "backend-go/internal/app/model"

func (d *Dao) UploadFile(upload model.Upload) (err error) {
	return d.db.Create(&upload).Error
}
