package main

import (
	"student-grade-api/config"
	"student-grade-api/models"
	"student-grade-api/handlers"
	"student-grade-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// =========================
	// Database Connection
	// =========================
	config.ConnectDB()

	// Auto migrate tables
	config.DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Enrollment{},
		&models.Grade{},
	)

	r := gin.Default()

	// =========================
	// Public Routes
	// =========================
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// =========================
	// Protected Routes (JWT Required)
	// =========================
	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		// Any logged-in user
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			role, _ := c.Get("role")

			c.JSON(200, gin.H{
				"user_id": userID,
				"role":    role,
				"message": "You accessed protected route",
			})
		})

		// =========================
		// Admin Routes
		// =========================
		adminGroup := protected.Group("/admin")
		adminGroup.Use(middleware.AuthorizeRole("admin"))
		{
			adminGroup.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Welcome Admin",
				})
			})

			adminGroup.POST("/courses", handlers.CreateCourse)
			adminGroup.GET("/courses", handlers.GetCourses)
			adminGroup.POST("/enroll", handlers.EnrollStudent)
		}

		// =========================
		// Teacher Routes
		// =========================
		teacherGroup := protected.Group("/teacher")
		teacherGroup.Use(middleware.AuthorizeRole("teacher"))
		{
			teacherGroup.POST("/assign-grade", handlers.AssignGrade)
		}

		// =========================
		// Student Routes
		// =========================
		studentGroup := protected.Group("/student")
		studentGroup.Use(middleware.AuthorizeRole("student"))
		{
			studentGroup.GET("/grades", handlers.ViewStudentGrades)
		}
	}

	// Start server
	r.Run(":8080")
}