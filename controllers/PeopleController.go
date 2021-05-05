package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/ivaroliRU/peopleAPI/business"
	"net/http"
	"strconv"
)

// GET api/people/
func GetPeople(c echo.Context) error {
	people := business.GetPeople()
	return c.JSON(http.StatusOK, people)
}

// POST api/people/
func AddPerson(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// GET api/people/:id
func GetPersonByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		return c.String(http.StatusNotFound, "")
	}

	person := business.GetPersonByID(int32(id))

	if person == nil {
		return c.String(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusOK, person)
}

// DELETE api/people/:id
func DeletePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// PUT api/people/:id
func UpdatePersonByID(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}