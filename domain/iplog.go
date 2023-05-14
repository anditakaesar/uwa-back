package domain

import "gorm.io/gorm"

type Iplog struct {
	gorm.Model
	Address string
	Count   int
}
