package model

import "gorm.io/gorm"

type Ens struct {
	gorm.Model
	Address string `json:"address"`
	Domain  string `json:"domain"`
}
