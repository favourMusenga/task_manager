package db

import (
	"time"

	"gorm.io/gorm"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
)

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Profile string

const (
	WorkProfile     Profile = "work"
	PersonalProfile Profile = "personal"
)

type Todo struct {
	gorm.Model
	Title       string
	Description *string
	DueDate     *time.Time
	Priority    Priority `gorm:"default:'low'"`
	Profile     Profile  `gorm:"default:'personal'"`
	Completed   bool     `gorm:"default:false"`
	CategoryID  *uint
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type FocusType string

const (
	Focus FocusType = "focus"
	Rest  FocusType = "rest"
)

type FocusTimer struct {
	gorm.Model
	FocusType  FocusType
	DuringTime uint // Duration in minutes
	TaskID     *uint
	Task       Todo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
