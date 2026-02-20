# 🏗 Design Document
## Student Grade Management System API

---

# 1️⃣ Overview

The Student Grade Management System is a RESTful API built using Go (Golang) to simulate a university academic portal.  
It supports multi-role access (Admin, Teacher, Student) and provides grade management with GPA calculation.

The system follows a layered architecture with clear separation of concerns.

---

# 2️⃣ Architecture Style

The application follows a **Layered Clean Architecture**:


Client
↓
Router (Gin)
↓
Middleware (Auth + Role)
↓
Handlers (Business Logic)
↓
ORM (GORM)
↓
PostgreSQL Database


Each layer has a specific responsibility:

- Router → HTTP request handling
- Middleware → Authentication & Authorization
- Handlers → Business logic
- Models → Database schema
- Config → Database connection

---

# 3️⃣ Design Principles Used

### ✅ Separation of Concerns
Each package handles a specific responsibility:
- `models/` → Database schema
- `handlers/` → API logic
- `middleware/` → Authentication & authorization
- `config/` → Database connection

### ✅ Role-Based Access Control (RBAC)
Authorization is implemented via middleware:
- JWT stores user role
- Middleware validates role before allowing access

### ✅ Stateless Authentication
JWT tokens ensure:
- No session storage needed
- Scalability
- Decoupled frontend-backend interaction

### ✅ Security
- Passwords hashed using bcrypt
- Token expiration implemented
- Protected routes require JWT validation

---

# 4️⃣ Database Design

## Entity Relationship Overview

Users
  ↳ Courses (TeacherID)
  ↳ Enrollments (StudentID)
       ↳ Grades (EnrollmentID)

---

## Table Relationships

### Users
- One Teacher → Many Courses
- One Student → Many Enrollments

### Courses
- Linked to Teacher
- Has multiple enrollments

### Enrollments
- Bridge table between Student and Course
- Used to track student-course relationship

### Grades
- Linked to Enrollment
- Stores marks and grade letter

---

# 5️⃣ Authentication Flow

1. User registers (password hashed)
2. User logs in
3. JWT token generated containing:
   - user_id
   - role
4. Client sends token in:
   Authorization: Bearer <token>
5. Middleware validates token
6. Role middleware restricts access

---

# 6️⃣ GPA Calculation Design

GPA is calculated dynamically:

For each grade:
- Marks converted to GPA scale
- Total GPA divided by total courses

Conversion scale:

| Marks | GPA |
|--------|------|
| 90+ | 4.0 |
| 80–89 | 3.5 |
| 70–79 | 3.0 |
| 60–69 | 2.5 |
| <60 | 0.0 |

Average marks are also calculated per student.

---

# 7️⃣ API Flow Example

Example: Grade Assignment

1. Teacher logs in
2. JWT token issued
3. Teacher calls:
   POST /protected/teacher/assign-grade
4. Middleware:
   - Validates JWT
   - Validates role = teacher
5. Grade stored in database
6. Letter grade auto-generated

---

# 8️⃣ Scalability Considerations

Future improvements may include:

- Pagination for large datasets
- Course-specific teacher validation
- Docker containerization
- Unit testing
- Swagger/OpenAPI documentation
- Caching layer (Redis)
- Microservice refactoring

---

# 9️⃣ Why Gin + GORM?

### Gin:
- Lightweight
- High performance
- Clean routing
- Middleware support

### GORM:
- Auto migration
- Relationship support
- Clean model-based schema
- Simplifies SQL operations

---

# 🔟 Conclusion

The system successfully implements:

- Secure authentication
- Role-based authorization
- Academic performance management
- GPA calculation
- Structured backend architecture

The project is scalable and production-ready with proper security and modular design.