package model

import (
	"gorm.io/gorm"
)

type Listitemdata struct {
	Model  `bson:"-"`
	MID_   string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Title  string `bson:"title" json:"title" xml:"title"`
	Itemid int64  `bson:"itemid" json:"itemid" xml:"itemid"`
}

func (Listitemdata) TableName() string {
	return "listitemdata"
}
func (m *Listitemdata) PreloadListitemdata(db *gorm.DB) *gorm.DB {
	return db
}
