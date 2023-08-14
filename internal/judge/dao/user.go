package dao

import (
	"time"
)

// GetUserResourceBeforeList 获取活跃时间在之前的地址
func (d *Dao) GetUserResourceBeforeList(before time.Time) (userList []string, err error) {
	for k, v := range d.active {
		if v.Before(before) {
			userList = append(userList, k)
		}
	}
	return
}

func (d *Dao) UpdateUserResourceTime(address string) (err error) {
	d.active[address] = time.Now()
	return
}
