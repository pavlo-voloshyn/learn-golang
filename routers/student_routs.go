package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pavlo-voloshyn/init/controllers"
	"github.com/pavlo-voloshyn/init/services"
)

var (
	student_service    = services.NewStudentService()
	student_controller = controllers.NewStudentController(student_service)
)

func SetStudentRouts(rg *gin.RouterGroup) {
	rg.GET("/student/:id", student_controller.GetById)
	rg.GET("/student", student_controller.GetAll)
	rg.POST("/student", student_controller.Create)
	rg.DELETE("/student", student_controller.Delete)
}
