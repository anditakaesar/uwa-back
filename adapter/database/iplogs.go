package database

import (
	"github.com/anditakaesar/uwa-back/domain"
	"gorm.io/gorm"
)

type IpLogRepositoryInterface interface {
	GetIplogByAddress(address string) (*domain.Iplog, error)
	UpdateCounter(address string) (*domain.Iplog, error)
}

type IpLogRepository struct {
	DB DatabaseInterface
}

func NewIpLogRepository(db DatabaseInterface) IpLogRepositoryInterface {
	return &IpLogRepository{DB: db}
}

func (a *IpLogRepository) GetIplogByAddress(address string) (*domain.Iplog, error) {
	db := a.DB.Get()
	var iplog domain.Iplog

	db.Where("address = ?", address).First(&iplog)

	return &iplog, nil
}

func (a *IpLogRepository) UpdateCounter(address string) (*domain.Iplog, error) {
	db := a.DB.Get()

	currentIplog := domain.Iplog{}
	result := db.Where("address = ?", address).First(&currentIplog)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	if currentIplog.ID == 0 {
		currentIplog.Address = address
		currentIplog.Count = 1
		result := db.Create(&currentIplog)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	currentIplog.Count += 1
	result = db.Save(&currentIplog)
	return &currentIplog, result.Error
}
