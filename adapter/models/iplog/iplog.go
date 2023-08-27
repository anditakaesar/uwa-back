package iplog

import (
	"github.com/anditakaesar/uwa-back/adapter/database"
	"gorm.io/gorm"
)

type Iplog struct {
	gorm.Model
	Address string
	Count   int
}

type IplogModelInterface interface {
	GetIplogByAddress(address string) (Iplog, error)
	UpdateCounter(address string) (Iplog, error)
}

type IplogModel struct {
	DB database.DatabaseInterface
}

func NewIplogModel(db database.DatabaseInterface) IplogModelInterface {
	return &IplogModel{DB: db}
}

func (a *IplogModel) GetIplogByAddress(address string) (Iplog, error) {
	db := a.DB.Get()
	var iplog Iplog

	db.Where("address = ?", address).First(&iplog)

	return iplog, nil
}

func (a *IplogModel) UpdateCounter(address string) (Iplog, error) {
	db := a.DB.Get()

	currentIplog := Iplog{}
	result := db.Where("address = ?", address).First(&currentIplog)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return currentIplog, result.Error
	}

	if currentIplog.ID == 0 {
		currentIplog.Address = address
		currentIplog.Count = 1
		result := db.Create(&currentIplog)
		if result.Error != nil {
			return currentIplog, result.Error
		}
	}

	currentIplog.Count += 1
	result = db.Save(&currentIplog)
	return currentIplog, result.Error
}
