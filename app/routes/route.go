package routes

import (
	"backend/app/middlewares"
	"backend/controllers/doctors"
	"backend/controllers/patients"
	"backend/controllers/schedules"
	"backend/controllers/visitors"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	DoctorController   doctors.DoctorController
	PatientController  patients.PatientController
	ScheduleController schedules.ScheduleController
	VisitorController  visitors.VisitorController
}

func (cl *ControllerList) DoctorRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	ev1 := e.Group("/api/v1/doctor")
	ev1.POST("/login", cl.DoctorController.Login)
	ev1.PUT("/update", cl.DoctorController.Update)
	ev1.POST("/register", cl.DoctorController.Register)
	ev1.GET("/getAll", cl.DoctorController.GetAll)
	// ev1.POST("/addSchedule", cl.DoctorController.AddSchedule)

}

func (cl *ControllerList) PatientRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/patient")
	ev1.POST("/login", cl.PatientController.Login)
	ev1.PUT("/update", cl.PatientController.Update)
	ev1.POST("/register", cl.PatientController.Register)
}

func (cl *ControllerList) ScheduleRouteRegister(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/schedule", middleware.JWTWithConfig(cl.JWTMiddleware))
	ev1.POST("/add", cl.ScheduleController.Add, middlewares.RoleValidation("doctor"))
	ev1.PUT("/update", cl.ScheduleController.Modif, middlewares.RoleValidation("doctor"))
	ev1.POST("/show", cl.ScheduleController.Show)
	ev1.DELETE("/delete", cl.ScheduleController.Remove, middlewares.RoleValidation("doctor"))
	ev1.GET("/getAll", cl.ScheduleController.GetAll)
	ev1.GET("/getAllDoctor", cl.ScheduleController.GetAllInOneDoctor)
	// ev1.POST("/insertDoctor", cl.ScheduleController.InsertDoctor)
}

func (cl *ControllerList) VisitorRoute(e *echo.Echo, ctx time.Duration) {
	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.BodyDump(middlewares.Log))
	ev1 := e.Group("/api/v1/visitor", middleware.JWTWithConfig(cl.JWTMiddleware))
	ev1.POST("/add", cl.VisitorController.AddVisitor, middlewares.RoleValidation("patient"))
	ev1.POST("/show", cl.VisitorController.ShowVisitor)
	ev1.DELETE("/delete", cl.VisitorController.DeleteVisitor)
	ev1.PUT("/update", cl.VisitorController.UpdateVisitor)
	ev1.DELETE("/cancel", cl.VisitorController.CancelVisitor, middlewares.RoleValidation("patient"))
	ev1.DELETE("/dontCome", cl.VisitorController.DontCome, middlewares.RoleValidation("doctor"))
	ev1.GET("/getAllPatient", cl.VisitorController.FetchAllPatient)
	ev1.GET("/showAllLog", cl.VisitorController.GetLogPatient, middlewares.RoleValidation("doctor"))
	ev1.GET("/ShowDetailSchedule", cl.VisitorController.GetScheduleDetails)

}
