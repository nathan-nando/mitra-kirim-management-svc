package model

import "time"

type LogVisitor struct {
	ID   string    `gorm:"column:id;primaryKey"`
	Ip   string    `gorm:"column:ip"`
	Date time.Time `gorm:"column:date"`
}

func (a *LogVisitor) TableName() string {
	return "log_visitor"
}
