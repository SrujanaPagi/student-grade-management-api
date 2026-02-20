package handlers

import (
	"net/http"
	"student-grade-api/config"
	"student-grade-api/models"

	"github.com/gin-gonic/gin"
)

type EnrollInput struct {
	StudentID uint `json:"student_id"`
	CourseID  uint `json:"course_id"`
}

func EnrollStudent(c *gin.Context) {
	var input EnrollInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	enrollment := models.Enrollment{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
	}

	config.DB.Create(&enrollment)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Student enrolled successfully",
		"enrollment": enrollment,
	})
}