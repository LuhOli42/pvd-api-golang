package model

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Categories struct {
	ID          uint   `json:"id" gorm:"primaryKey" valid:"-"`
	Description string `json:""description" gorm:"type:varchar(255)" valid:"notnull"`
}
