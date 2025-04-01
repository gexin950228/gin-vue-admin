package global

import "time"

type GVA_MODEL struct {
	ID        uint      `gorm:"primary" json:"ID"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt time.Time `gorm:"index" json:"-"` // 删除时间
}
