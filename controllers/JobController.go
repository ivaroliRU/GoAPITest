package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/ivaroliRU/peopleAPI/business"
	"net/http"
	"strconv"
)

// GET api/jobs/
func GetJobs(c echo.Context) error {
	jobs := business.GetJobs()
	return c.JSON(http.StatusOK, jobs)
}

// POST api/jobs/
func AddJob(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/jobs/:id
func GetJobByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusNotFound, "")
	}

	job := business.GetJobsByID(int32(id))

	if job == nil {
		return c.String(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusOK, job)
}

// GET api/jobs/:id/people
func GetPeopleWithJob(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusNotFound, "")
	}

	people := business.GetPeopleWithJob(int32(id))
	return c.JSON(http.StatusOK, people)
}

// PUT api/people/:id
func UpdateJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}