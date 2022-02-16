package models

import (
	"goblog/pkg/types"
)

type BaseModel struct {
	ID uint64 `gorm:"primary_key"`
}

func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
