package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	EnrollmentID uint
	Marks        float64
	GradeLetter  string
}