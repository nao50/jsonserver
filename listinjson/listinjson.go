package listinjson

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ListInJson struct {
	Json []*Json `json:"json"`
}

type Json struct {
	JsonId   int    `json:"jsonId"`
	JsonName string `json:"jsonName"`
}

var (
	jsons         = map[int]*Json{}
	jsonssequence = 1
)

func CreateListInJson(c echo.Context) error {
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

func GetListInJsons(c echo.Context) error {
	js := []*Json{}
	for _, j := range jsons {
		js = append(js, j)
	}
	jsonlist := &ListInJson{
		Json: js,
	}
	return c.JSON(http.StatusOK, jsonlist)
}

func UpdateListInJson(c echo.Context) error {
	j := new(Json)
	if err := c.Bind(j); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("jsonid"))
	jsons[id].JsonName = j.JsonName
	return c.JSON(http.StatusOK, jsons[id])
}

func DeleteListInJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("jsonid"))
	delete(jsons, id)
	return c.NoContent(http.StatusNoContent)
}
