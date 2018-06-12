package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naoyamaguchi/jsonserver/authentication"
	"github.com/naoyamaguchi/jsonserver/json"
	"github.com/naoyamaguchi/jsonserver/listinjson"
	"github.com/naoyamaguchi/jsonserver/nestedjson"
	"github.com/naoyamaguchi/jsonserver/niljson"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// authentication
	e.POST("/auth", authentication.Authentication)

	// must authenticated group
	r := e.Group("/")
	r.Use(middleware.JWTWithConfig(authentication.JWTConfiguration()))

	// json
	r.POST("jsons", json.CreateJson)
	r.GET("jsons", json.GetJsons)
	r.GET("jsons/:jsonid", json.GetJson)
	r.PUT("jsons/:jsonid", json.UpdateJson)
	r.DELETE("jsons/:jsonid", json.DeleteJson)

	// nestedjson
	r.POST("nestedjsons", nestedjson.CreateNestedJson)
	r.GET("nestedjsons", nestedjson.GetNestedJsons)
	r.GET("nestedjsons/:jsonid", nestedjson.GetNestedJson)
	r.PUT("nestedjsons/:jsonid", nestedjson.UpdateNestedJson)
	r.DELETE("nestedjsons/:jsonid", nestedjson.DeleteNestedJson)

	// listinjson
	r.POST("listinjsons", listinjson.CreateListInJson)
	r.GET("listinjsons", listinjson.GetListInJsons)
	r.PUT("listinjsons/:jsonid", listinjson.UpdateListInJson)
	r.DELETE("listinjsons/:jsonid", listinjson.DeleteListInJson)

	// jsoninlist

	// niljson
	r.POST("niljsons", niljson.CreateNilJson)
	r.GET("niljsons", niljson.GetNilJsons)

	e.Logger.Fatal(e.Start(":8080"))
}
