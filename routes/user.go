package routes

import (
	"planeta/controllers"
)

// Get a UserController instance
var p = controllers.PlanetController{}

func init() {

	R.GET("/teste/v1/planets/:planeta", p.GetPlanet)
	R.POST("/teste/v1/planets/", p.PostPlanet)
	R.DELETE("/teste/v1/planets/:planeta", p.DeletePlanet)
	R.GET("/teste/v1/planets/", p.ListPlanet)
}
