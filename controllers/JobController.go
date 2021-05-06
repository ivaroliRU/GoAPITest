package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/ivaroliRU/peopleAPI/business"
	"github.com/ivaroliRU/peopleAPI/data"
	"github.com/ivaroliRU/peopleAPI/models"
	"net/http"
	"strconv"
)

var jobBusiness business.JobsBusiness = business.JobsBusiness {
	Data: &data.Data{},
}

// GET api/jobs/
func GetJobs(c echo.Context) error {
	jobs := jobBusiness.GetJobs()
	return c.JSON(http.StatusOK, jobs)
}

// POST api/jobs/
func AddJob(c echo.Context) error {
	j := new(models.Job)
	err := c.Bind(j)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	jobBusiness.AddJob(*j)

	return c.JSON(http.StatusOK, j)
}

// GET api/jobs/:id
func GetJobByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusNotFound, "")
	}

	job := jobBusiness.GetJobsByID(int32(id))

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

	people := jobBusiness.GetPeopleWithJob(int32(id))
	return c.JSON(http.StatusOK, people)
}

// PUT api/people/:id
func UpdateJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}