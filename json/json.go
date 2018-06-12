package json

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Json struct {
	JsonId   int    `json:"jsonId"`
	JsonName string `json:"jsonName"`
}

var (
	jsons         = map[int]*Json{}
	jsonssequence = 2
)

// initialization
func init() {
	j := &Json{
		JsonId:   1,
		JsonName: "nananao",
	}
	jsons[j.JsonId] = j
}

// POST /json
func CreateJson(c echo.Context) error {
	j := &Json{
		JsonId: jsonssequence,
	}
	if err := c.Bind(j); err != nil {
		return err
	}
	jsons[j.JsonId] = j
	jsonssequence++
	return c.JSON(http.StatusCreated, j)
}

// GET /jsons
func GetJsons(c echo.Context) error {
	js := []*Json{}
	for _, j := range jsons {
		js = append(js, j)
	}
	return c.JSON(http.StatusOK, js)
}

// GET /json/:jsonid
func GetJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("jsonid"))
	return c.JSON(http.StatusOK, jsons[id])
}

// PUT /json/:jsonid
func UpdateJson(c echo.Context) error {
	j := new(Json)
	if err := c.Bind(j); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("jsonid"))
	jsons[id].JsonName = j.JsonName
	return c.JSON(http.StatusOK, jsons[id])
}

// DELETE /json/:jsonid
func DeleteJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("jsonid"))
	delete(jsons, id)
	return c.NoContent(http.StatusNoContent)
}
