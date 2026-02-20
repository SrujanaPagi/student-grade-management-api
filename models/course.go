package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name      string
	TeacherID uint
}