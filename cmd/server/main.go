package main

import (
	"log"

	"github.com/gravestench/jh-weather-exercise/pkg/configManager"
	"github.com/gravestench/jh-weather-exercise/pkg/ginWebServer"
	"github.com/gravestench/jh-weather-exercise/pkg/openWeatherMap"
)

const (
	configDirectory = "~/jh_weather_exercise" // our config files will go here
)

func main() {
	// the main "services" of our application
	var (
		svcConfigManager configManager.Service  // manages configs for other services
		svcWebServer     ginWebServer.Service   // is the web server and router
		svcWeather       openWeatherMap.Service // exposes the Open Weather Map API
	)

	svcConfigManager.RootDirectory = configDirectory

	// first, for services that have config files, we want to
	// init them with the config manager service, which will load
	// the file and set defaults if necessary.
	for _, candidate := range []configManager.Configurable{
		&svcWebServer,
		&svcWeather,
	} {
		if err := svcConfigManager.InitConfiguration(candidate); err != nil {
			log.Fatalf("initializing config: %v", err)
		}
	}

	// now we need to pass our route-initializers to the server/router.
	// if we had other services that implemented the GinRouteInitializer
	// interface then we would put them here
	for _, candidate := range []ginWebServer.GinRouteInitializer{
		&svcConfigManager,
		&svcWeather,
	} {
		svcWebServer.InitializeRoutes(candidate)
	}

	// finally, kick of the web server
	log.Printf("serving on port %d", svcWebServer.Port)
	if err := svcWebServer.Serve(); err != nil {
		log.Fatalf("web server stopped serving: %v", err)
	}
}
