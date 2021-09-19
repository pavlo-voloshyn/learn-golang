package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/entities"
	"github.com/pavlo-voloshyn/init/services"
)

type StudentController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type studentController struct {
	studentService services.StudentService
}

func NewStudentController(st services.StudentService) StudentController {
	return &studentController{
		studentService: st,
	}
}

func (s *studentController) GetById(c *gin.Context) {
	c.JSON(200, s.studentService.GetById(c.GetInt("id")))
}

func (s *studentController) GetAll(c *gin.Context) {
	c.JSON(200, s.studentService.GetAll())
}

func (s *studentController) Create(c *gin.Context) {
	var student entities.Student
	c.Bind(&student)
	s.studentService.Create(student)
	c.AbortWithStatus(http.StatusOK)
}

func (s *studentController) Delete(c *gin.Context) {
	if err := s.studentService.Delete(c.GetInt("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	c.AbortWithStatus(http.StatusOK)

}
