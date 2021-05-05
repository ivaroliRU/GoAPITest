package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GET api/jobs/
func index(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// POST api/jobs/
func addJob(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/jobs/:id
func getJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/jobs/:id/people
func getPeopleWithJob(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PUT api/people/:id
func updateJobByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}