package nestedjson

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type NestedJson struct {
	NestedJsonId   int           `json:"nestedJsonId"`
	NestedJsonName string        `json:"nestedJsonName"`
	Configuration  Configuration `json:"configuration"`
}

type Configuration struct {
	// ConfigurationId   string `json:"configurationId"`
	ConfigurationName string `json:"configurationName"`
	Configuration     string `json:"configuration"`
}

var (
	nestedjsons         = map[int]*NestedJson{}
	nestedjsonssequence = 2
)

// initialization
func init() {
	j := &NestedJson{
		NestedJsonId:   1,
		NestedJsonName: "nestedFugaJson",
		Configuration: Configuration{
			ConfigurationName: "nestedFugaConfiguration",
			Configuration:     "Middle",
		},
	}
	nestedjsons[j.NestedJsonId] = j
}

// POST /nestedjson
func CreateNestedJson(c echo.Context) error {
	j := &NestedJson{
		NestedJsonId: nestedjsonssequence,
	}
	if err := c.Bind(j); err != nil {
		return err
	}
	nestedjsons[j.NestedJsonId] = j
	nestedjsonssequence++
	return c.JSON(http.StatusCreated, j)
}

// GET /nestedjson
func GetNestedJsons(c echo.Context) error {
	js := []*NestedJson{}
	for _, j := range nestedjsons {
		js = append(js, j)
	}
	return c.JSON(http.StatusOK, js)
}

// GET /nestedjson/:jsonid
func GetNestedJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("jsonid"))
	return c.JSON(http.StatusOK, nestedjsons[id])
}

// PUT /nestedjson/:jsonid
func UpdateNestedJson(c echo.Context) error {
	j := new(NestedJson)
	if err := c.Bind(j); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("jsonid"))
	// nestedjsons[id] = j
	nestedjsons[id].NestedJsonName = j.NestedJsonName
	nestedjsons[id].Configuration = j.Configuration
	return c.JSON(http.StatusOK, nestedjsons[id])
}

// DELETE /nestedjson/:jsonid
func DeleteNestedJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("jsonid"))
	delete(nestedjsons, id)
	return c.NoContent(http.StatusNoContent)
}
