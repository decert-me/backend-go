package model

type SystemSetting struct {
	ID    uint   `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id" form:"id"`
	Key   string `gorm:"column:key;type:varchar(255);comment:键" json:"key" form:"key"`
	Value string `gorm:"column:value;type:varchar(255);comment:值" json:"value" form:"value"`
	Desc  string `gorm:"column:desc;type:varchar(255);comment:描述" json:"desc" form:"desc"`
}

func (SystemSetting) TableName() string {
	return "admin_system_setting"
}
