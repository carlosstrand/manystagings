package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Environment struct {
	Base
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (e *Environment) BeforeCreate(tx *gorm.DB) (err error) {
	if e.Namespace == "" {
		e.Namespace = slug.Make(e.Name)
	}
	return
}
