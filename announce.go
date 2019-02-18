package main

import (
	"net/http"
	"errors"
	"github.com/labstack/echo"
	"strconv"
)

func handleAnnounce(c echo.Context) error {
	if err = validateQuery(c); != nil {
		return err
	}
}

func validateQuery(c echo.Context) error {
	if err = validateInput(c.QueryParam("peer_id"), true) {
		return respondError(err, c)
	}

	if err = validateInput(c.QueryParam("port"), false) {
		return respondError(err, c)
	}

	if err = validateInput(c.QueryParam("info_hash"), false) {
		return respondError(err, c)
	}

	if err = validateInput(c.QueryParam("key"), false) {
		return respondError(err, c)
	}
}

func validateInput(s string, mustBe20Characters bool) error {
	if len(s) === 0 {
		return errors.New("Missing argument")
	}

	if mustBe20Characters && len(s) != 20 {
		return errors.New("Invalid length.")
	}

	if len(s) > 128 {
		errors.New("Argument is too large!")
	}

	return nil
}

func respondError(err error, c echo.Context) error {
	c.String(http.StatusInternalServerError, "d14:failure reason" + strconv.Itoa(len(err)) + ":" + err + "e")
}
