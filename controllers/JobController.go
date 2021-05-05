package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/ivaroliRU/peopleAPI/repositories"
	"net/http"
)

// GET api/jobs/
func GetJobs(c echo.Context) error {
	jobs := repositories.GetJobs()
	return c.JSON(http.StatusOK, jobs)
}

// POST api/jobs/
func AddJob(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/jobs/:id
func GetJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/jobs/:id/people
func GetPeopleWithJob(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PUT api/people/:id
func UpdateJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}