package handlers

import (
	"net/http"
	"student-grade-api/config"
	"student-grade-api/models"

	"github.com/gin-gonic/gin"
)

type CreateCourseInput struct {
	Name      string `json:"name"`
	TeacherID uint   `json:"teacher_id"`
}

func CreateCourse(c *gin.Context) {
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	course := models.Course{
		Name:      input.Name,
		TeacherID: input.TeacherID,
	}

	config.DB.Create(&course)

	c.JSON(http.StatusOK, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	config.DB.Find(&courses)

	c.JSON(http.StatusOK, courses)
}