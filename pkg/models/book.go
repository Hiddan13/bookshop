package models

import (
	"time"

	_ "github.com/Hiddan13/bookshop/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Book struct {
	gorm.Model
	Name      string `gorm:"" json:"namebook"`
	Author    Author `gorm:"" json:"author"`
	Publisher string `gorm:"" json:"publisher"`
}
type Author struct {
	gorm.Model
	Name           string `gorm:"" json:"nameauthor"`
	LastName       string `gorm:"" json:"lastnameauthor"`
	PublishedBooks int    `gorm:"" json:"publishedbooks"`
}
