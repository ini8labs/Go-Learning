<h1 style="text-align:center;">SWAGGER for Gin</h1>

## SWAGGER

- OPEN Source API Documentation framework to help developers in design, document, and consume RESTful web services.

- Swagger reads an API and extract in the form of interactive UI called as "Swagger UI".

- Swagger Ul offers HTML view of API with JSON Support.

- Most popular tool for generating interactive documentation from API.

## Need for documentation of API

- What are the endpoint URLs?

- What method to be used?

- What data to be passed?

- What are the mandatory and optional params to the endpoint?

- Is authentication required?

## GitHub repository of gin-swagger

- [gin-wagger GitHub Repo](https://github.com/swaggo/gin-swagger)

-     https://github.com/swaggo/gin-swagger

## SWAG installation in Golang

- Swag is middleware that helps to automatically generate RESTful API documentation with Swagger 2.0 for Go directly from source code using annotations.

### Run in terminal  

-     go get -u github.com/swaggo/swag/cmd/swag

-     go install  github.com/swaggo/swag/cmd/swag@latest

#### Note

- To check if swag is installed run `swag -h`

## Add Library

-     go get -u github.com/swaggo/files
-     go get -u github.com/swaggo/gin-swagger

## swag init

- To create and update the necessary swagger documentation.

- Run `swag init` before running *main.go*

- This step needs to be performed for each iteration.



## Open SwaggerUI in browser

-     http://localhost:8080/swagger/index.html

## OpenAPI Specification

- Formerly known as the Swagger Specification, is the world’s standard for defining RESTful interfaces.
