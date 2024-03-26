# Weather Service Coding Exercise

## Instructions
> Write an http server that uses the Open Weather API that exposes an endpoint that takes in lat/long
coordinates. This endpoint should return what the weather condition is outside in that area (snow, rain,
etc), whether it’s hot, cold, or moderate outside (use your own discretion on what temperature equates to
each type).
> 
> The API can be found here:https://openweathermap.org/api. Even though most of the API calls found on
OpenWeather aren’t free, you should be able to use the free “current weather data” API call for this
project. First, sign-up for an account, which shouldn’t require credit card or payment information. Once
you’ve created an account, use https://openweathermap.org/faq to get your API Key setup to start using
the API.
>
> you’ve coded your project, add it to a publicly accessible Github repository and share it
with the team. Additionally, please don’t add your API Key to the project. Each member of the
team reviewing your code has their own key to use for testing your project.

## My Implementation
The code I've written for this exercise uses the "standard"
Golang project structure. Stuff that I treat like library 
components are put inside of `pkg`, and applications each
have a subdir inside of `cmd`.

There are three "services" that make up the application:
* **Config File Manager** - Initializes configs in a standard way, but doesn't know anything about the actual config files. The file defaults and ingestion is done by other delegate services which integrate with this one by implementing an integration interface.
* **Web Server/Router** - A combined HTTP web server and a gin web router. Other services integrate with this one by implementing an interface that allows them to be delegates for initializing their API routes.
* **Open Weather Map API** - Exposes api endpoints that internally use the underlying Open Weather Map API.

These three services are all set uptogether in the `main` func of the 
application in `cmd/server/main.go`. 

## Running the Example Application

You can run the application by executing this command:
```bash
go run github.com/gravestench/jh-weather-exercise/cmd/server@main
```

The first time the application runs, it should create a subdirectory in the current user's 
home directory `~/jh_weather_exercise`. This is where the config files will be saved/loaded from.
The program should exit with an error message telling you to set up your API key, and it should 
print the path of the config file you need to edit, and also where to get the api key.

The `openWeatherMap` service only has two routes:
* `weather/current`
* `weather/current/describe`

For example, once the app is running you can hit the exposed API endpoint `http://localhost:8080/weather/current`

Both of these endpoints take the same JSON payload. Here is an example of the payload object, with 
latitude and longitude for Fresno, CA:
```json
{
  "Latitude": 36.73,
  "Longitude": 119.78
}
```

Once the app is running, you can use a tool like Postman to test the endpoints.

The output of the `weather/current/describe` endpoint yields a response like this:
```text
"It's cool, feels colder than it actually is, and dry, with clear skies."
```

## Bonus
There's also an api route defined by the config manager:
* `http://localhost:8080/config/paths`

This was just added to show how easily it is to add other services to the API of our example app. 
This route just prints the filepaths of the config files. 