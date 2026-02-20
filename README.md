# 🎓 Student Grade Management System API

A production-ready RESTful API built using **Go (Golang)** that manages student academic performance, including course management, enrollment, grade assignment, GPA calculation, and role-based access control.

This system simulates university portals such as Canvas or Blackboard and supports multiple user roles: **Admin, Teacher, and Student**.

---

## 🚀 Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Token)
- **Password Security:** bcrypt
- **Architecture Style:** Layered Clean Architecture

---

## 🏗 System Architecture


Client (Thunder Client / Postman)
↓
Gin Router
↓
Middleware (JWT + Role Authorization)
↓
Handlers
↓
GORM ORM
↓
PostgreSQL Database

---

## 👥 User Roles & Permissions

| Feature | Admin | Teacher | Student |
|----------|--------|----------|-----------|
| Register/Login | ✅ | ✅ | ✅ |
| Create Course | ✅ | ❌ | ❌ |
| View Courses | ✅ | ❌ | ❌ |
| Enroll Student | ✅ | ❌ | ❌ |
| Assign Grade | ❌ | ✅ | ❌ |
| View Own Grades | ❌ | ❌ | ✅ |
| GPA Calculation | ❌ | ❌ | ✅ |

---

## 🗄 Database Schema

### Users
- ID
- Name
- Email
- Password (bcrypt hashed)
- Role (admin / teacher / student)

### Courses
- ID
- Name
- TeacherID (Foreign Key)

### Enrollments
- ID
- StudentID (Foreign Key)
- CourseID (Foreign Key)

### Grades
- ID
- EnrollmentID (Foreign Key)
- Marks
- GradeLetter

---

## 🔐 Authentication & Authorization

- Passwords are securely hashed using **bcrypt**
- JWT tokens are generated during login
- JWT includes:
  - `user_id`
  - `role`
- Middleware validates:
  - Token authenticity
  - Role-based permissions

---

## 🎓 Features Implemented

### ✅ Admin
- Create courses
- Enroll students into courses
- View course list

### ✅ Teacher
- Assign grades to enrolled students
- Automatic letter grade generation

### ✅ Student
- View personal grades
- GPA calculation
- Average marks calculation
- Total enrolled courses count

---

## 📊 GPA Calculation Logic

| Marks | GPA |
|--------|------|
| 90+ | 4.0 |
| 80-89 | 3.5 |
| 70-79 | 3.0 |
| 60-69 | 2.5 |
| <60 | 0.0 |

GPA is calculated dynamically based on all assigned grades.

---

## 🧪 API Endpoints

### Public Routes

POST /register
POST /login


### Protected Routes (JWT Required)

#### Admin

GET /protected/admin/dashboard
POST /protected/admin/courses
GET /protected/admin/courses
POST /protected/admin/enroll


#### Teacher

POST /protected/teacher/assign-grade


#### Student

GET /protected/student/grades


---

## ⚙️ Setup Instructions

### 1️⃣ Clone Repository


git clone https://github.com/SrujanaPagi/student-grade-management-api.git

cd student-grade-management-api


### 2️⃣ Install Dependencies


go mod tidy


### 3️⃣ Setup PostgreSQL

Create database:


CREATE DATABASE grades;


Update database credentials in:


config/database.go


### 4️⃣ Run Application


go run cmd/main.go


Server runs on:


http://localhost:8080


---

## 🧠 Design Decisions

- Used Gin for lightweight and fast routing.
- Used GORM for ORM and automatic migrations.
- Implemented layered architecture for scalability.
- JWT chosen for stateless authentication.
- Role-based middleware ensures clean authorization logic.

---

## 📈 Future Improvements

- Swagger/OpenAPI documentation
- Pagination for course listing
- Docker containerization
- Unit testing
- Refresh token mechanism
- Course-specific grade validation
- Admin dashboard analytics

---

## 📌 Project Status

✅ Fully functional  
✅ Role-based access implemented  
✅ Academic performance tracking completed  
✅ Ready for evaluation  

---

## 👩‍💻 Author

Srujana Pagi  
Student Grade Management System – Capstone Project  