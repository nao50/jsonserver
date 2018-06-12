package niljson

import (
	"net/http"

	"github.com/labstack/echo"
)

type NilJson struct {
	NilJsonId       int     `json:"nilJsonId"`
	PureString      string  `json:"pureString"`
	PointerString   *string `json:"pointerString"`
	OmitemptyString string  `json:"omitemptyString,omitempty"`
}

var (
	jsons         = map[int]*NilJson{}
	jsonssequence = 1
)

// POST /json
func CreateNilJson(c echo.Context) error {
	j := &NilJson{
		NilJsonId: jsonssequence,
	}
	if err := c.Bind(j); err != nil {
		return err
	}
	jsons[j.NilJsonId] = j
	jsonssequence++
	return c.JSON(http.StatusCreated, j)
}

// GET /jsons
func GetNilJsons(c echo.Context) error {
	js := []*NilJson{}
	for _, j := range jsons {
		js = append(js, j)
	}
	return c.JSON(http.StatusOK, js)
}
