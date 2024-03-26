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

## Implementation
The code I've written for this exercise uses the "standard"
Golang project structure. Stuff that I treat like library 
components are put inside of `pkg`, and applications each
have a subdir inside of `cmd`.

There are three "services" that make up the application:
- **Config File Manager Service**: Handles configurations via an integration interface without direct file knowledge.
- **Web Server/Router Service**: Integrates an HTTP server with a gin web router, allowing route initialization through an interface.
- **Weather Service**: Utilizes the OpenWeather API to provide weather conditions via endpoints.

## Design Philosophy
The architecture of the project is built on the principles of modularity, testability, 
and clear division of responsibilities, making it straightforward to incorporate 
additional services. 

Key to my approach is the emphasis on these aspects:

1) **Config Service Flexibility**: This service empowers other components to 
manage their configuration data independently, supporting various data formats 
like JSON, YAML, or XML. This flexibility ensures the config manager needs only to 
know the file names for data access, making any format change seamless.

2) **Web Server and Router Integration**: This core component works closely with the 
config service, allowing other services to define their API routes through an 
integration interface. **A significant advantage of this method is avoiding 
the monolithic "god-file" approach for route initialization, which can lead to 
maintenance challenges.**

3) **Weather Service Integration**: By maintaining a clear boundary of responsibilities, the Weather Service efficiently interacts with the Config and Web Server/Router Services. This design ensures each service focuses on its core functionality while contributing to the application's overall capabilities.


### What I would do if this wasn't a coding exercise

If I were to create this application for myself, I would have opted to use my own micro-framework
([Service Mesh](https://github.com/gravestench/servicemesh)). It takes a similar approach in
terms of project organization and modularity.

## Running the Example Application

You can run the application by executing this command:
```bash
go run github.com/gravestench/jh-weather-exercise/cmd/example-app@main
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
{
    "Description": "It's cool, feels colder than it actually is, and dry, with clear skies."
}
```

## Bonus
There's also an api route defined by the config manager:
* `http://localhost:8080/config/paths`

This was just added to show how easily it is to add other services to the API of our example app. 
This route just prints the filepaths of the config files. The response json looks like this:
```json
{
    "Paths": [
        "C:\\Users\\dknuth\\jh_weather_exercise\\web_server.json",
        "C:\\Users\\dknuth\\jh_weather_exercise\\open_weather_map.json"
    ]
}
```