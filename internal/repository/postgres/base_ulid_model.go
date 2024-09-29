package postgres

import (
	"search-keyword-service/pkg/id"
	"time"

	"gorm.io/gorm"
)

type BaseULIDModel struct {
	ID        string    `gorm:"column:id;type:char(26);primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;autoUpdateTime" json:"updated_at"`
}

func NewBaseULIDModel() BaseULIDModel {
	return BaseULIDModel{
		ID: id.NewULID(),
	}
}

func (b *BaseULIDModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = id.NewULID()
	} else {
		if err := id.CheckIsULID(b.ID); err != nil {
			return err
		}
	}

	return nil
}
