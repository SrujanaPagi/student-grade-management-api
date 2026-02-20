package handlers

import (
	"net/http"
	"student-grade-api/config"
	"student-grade-api/models"

	"github.com/gin-gonic/gin"
)

type AssignGradeInput struct {
	EnrollmentID uint    `json:"enrollment_id"`
	Marks        float64 `json:"marks"`
}

func AssignGrade(c *gin.Context) {
	var input AssignGradeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Calculate letter grade
	letter := calculateLetterGrade(input.Marks)

	grade := models.Grade{
		EnrollmentID: input.EnrollmentID,
		Marks:        input.Marks,
		GradeLetter:  letter,
	}

	config.DB.Create(&grade)

	c.JSON(http.StatusOK, gin.H{
		"message": "Grade assigned successfully",
		"grade":   grade,
	})
}

func calculateLetterGrade(marks float64) string {
	if marks >= 90 {
		return "A"
	} else if marks >= 80 {
		return "B"
	} else if marks >= 70 {
		return "C"
	} else if marks >= 60 {
		return "D"
	}
	return "F"
}