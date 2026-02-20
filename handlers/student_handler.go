package handlers

import (
	"net/http"
	"student-grade-api/config"
	"student-grade-api/models"

	"github.com/gin-gonic/gin"
)

func ViewStudentGrades(c *gin.Context) {

	userIDInterface, _ := c.Get("user_id")
	userID := uint(userIDInterface.(float64))

	// Get enrollments for this student
	var enrollments []models.Enrollment
	config.DB.Where("student_id = ?", userID).Find(&enrollments)

	var grades []models.Grade
	var totalMarks float64
	var totalCourses int
	var totalGPA float64

	for _, enrollment := range enrollments {
		var grade models.Grade
		config.DB.Where("enrollment_id = ?", enrollment.ID).First(&grade)

		if grade.ID != 0 {
			grades = append(grades, grade)
			totalMarks += grade.Marks
			totalCourses++
			totalGPA += convertToGPA(grade.Marks)
		}
	}

	var averageMarks float64
	var gpa float64

	if totalCourses > 0 {
		averageMarks = totalMarks / float64(totalCourses)
		gpa = totalGPA / float64(totalCourses)
	}

	c.JSON(http.StatusOK, gin.H{
		"grades":        grades,
		"gpa":           gpa,
		"average_marks": averageMarks,
		"total_courses": totalCourses,
	})
}

func convertToGPA(marks float64) float64 {
	if marks >= 90 {
		return 4.0
	} else if marks >= 80 {
		return 3.5
	} else if marks >= 70 {
		return 3.0
	} else if marks >= 60 {
		return 2.5
	}
	return 0.0
}