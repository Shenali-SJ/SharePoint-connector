package model

import (
	"gorm.io/gorm"
)

type Requestdata struct {
	Model        `bson:"-"`
	MID_         string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Operation    string `bson:"operation" json:"operation" xml:"operation"`
	Siteurl      string `bson:"siteurl" json:"siteurl" xml:"siteurl"`
	Clientid     string `bson:"clientid" json:"clientid" xml:"clientid"`
	Clientsecret string `bson:"clientsecret" json:"clientsecret" xml:"clientsecret"`
	Listname     string `bson:"listname" json:"listname" xml:"listname"`
}

func (Requestdata) TableName() string {
	return "requestdata"
}
func (m *Requestdata) PreloadRequestdata(db *gorm.DB) *gorm.DB {
	return db
}
