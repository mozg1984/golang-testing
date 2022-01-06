package app

import (
	"github.com/mozg1984/golang-testing/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
