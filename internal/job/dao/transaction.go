package dao

import "backend-go/internal/app/model"

func (d *Dao) QueryWaitTransaction() (transHashList []model.Transaction, err error) {
	err = d.db.Where("status = 0").Find(&transHashList).Error
	return
}

func (d *Dao) UpdateTransactionStatus(tx *model.Transaction) (err error) {
	err = d.db.Model(&model.Transaction{}).
		Where("hash = ?", tx.Hash).
		Updates(map[string]interface{}{"status": tx.Status, "msg": tx.Msg}).Error
	return err
}

func (d *Dao) CreateTransaction(req *model.Transaction) (err error) {
	err = d.db.Model(&model.Transaction{}).Create(req).Error
	return

}

func (d *Dao) QueryTransactionByHash(hash string) (transHash model.Transaction, err error) {
	err = d.db.Where("hash", hash).First(&transHash).Error
	return
}
