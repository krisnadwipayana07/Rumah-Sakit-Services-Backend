package main

import (
	_middlewares "backend/app/middlewares"

	"backend/app/routes"

	_doctorUsecase "backend/business/doctors"
	_patientUsecase "backend/business/patients"
	_scheduleUsecase "backend/business/schedules"
	_visitorUsecase "backend/business/visitors"

	_doctorControllers "backend/controllers/doctors"
	_patientControllers "backend/controllers/patients"
	_scheduleControllers "backend/controllers/schedules"
	_visitorControllers "backend/controllers/visitors"

	_doctordb "backend/drivers/databases/doctors"
	_patientdb "backend/drivers/databases/patients"
	_scheduledb "backend/drivers/databases/schedules"
	_visitordb "backend/drivers/databases/visitors"

	_mysqlDriver "backend/drivers/mysql"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Echo Swagger Example API
// @version 1.0
// @description documentation swagger.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("This Service RUN on DEBUG Mode")
	}
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&_doctordb.Doctors{})
	db.AutoMigrate(&_patientdb.Patients{})
	db.AutoMigrate(&_scheduledb.Schedules{})
	db.AutoMigrate(&_visitordb.Visitors{})
	db.AutoMigrate(&_visitordb.VisitorsLog{})

}

func main() {
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	Conn := configDB.InitialDB()
	DBMigrate(Conn)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiredDuration: viper.GetInt(`jwt.expired`),
	}

	e := echo.New()

	// db, err := _mongoDB.Connect()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	doctorRepository := _doctordb.NewMysqlDoctorRepository(Conn)
	doctorUseCase := _doctorUsecase.NewDoctorUsecase(doctorRepository, &configJWT, timeoutContext)
	DoctorController := _doctorControllers.NewDoctorController(doctorUseCase)

	patientRepository := _patientdb.NewMysqlPatientRepository(Conn)
	patientUsecase := _patientUsecase.NewPatientsUsecase(patientRepository, &configJWT, timeoutContext)
	PatientController := _patientControllers.NewPatientController(patientUsecase)

	scheduleRepository := _scheduledb.NewMysqlSchedulesRepository(Conn)
	scheduleUsecase := _scheduleUsecase.NewSquedulesUsecase(scheduleRepository, timeoutContext)
	ScheduleController := _scheduleControllers.NewScheduleController(scheduleUsecase)

	visitorRepository := _visitordb.NewMysqlVisitorRepository(Conn)
	visitorUsecase := _visitorUsecase.NewVisitorUsecase(visitorRepository, timeoutContext)
	VisitorController := _visitorControllers.NewVisitorController(visitorUsecase)

	routesInit := routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		DoctorController:   *DoctorController,
		PatientController:  *PatientController,
		ScheduleController: *ScheduleController,
		VisitorController:  *VisitorController,
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routesInit.DoctorRouteRegister(e, timeoutContext)
	routesInit.PatientRouteRegister(e, timeoutContext)
	routesInit.ScheduleRouteRegister(e, timeoutContext)
	routesInit.VisitorRoute(e, timeoutContext)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
