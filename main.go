package main

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/ivaroliRU/peopleAPI/controllers"
)

func main() {
	// create echo instance
	e := echo.New()

	// simple health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	// People routing
	e.GET("/api/people", controllers.GetPeople)
	e.POST("/api/people", controllers.AddPerson)
	e.GET("/api/people/:id", controllers.GetPersonByID)
	e.PUT("/api/people/:id", controllers.UpdatePersonByID)
	e.DELETE("/api/people/:id", controllers.DeletePersonByID)

	// Job routing
	e.GET("/api/jobs", controllers.GetJobs)
	e.POST("/api/jobs", controllers.AddJob)
	e.GET("/api/jobs/:id", controllers.GetJobByID)
	e.GET("/api/jobs/:id/people", controllers.GetPeopleWithJob)
	e.PUT("/api/jobs/:id", controllers.UpdateJobByID)

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}