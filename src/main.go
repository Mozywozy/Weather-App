package main

import (
    "weather-ap/src/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("src/templates/*")
    router.Static("/static", "./src/static")

    router.GET("/", handler.GetWeather)

    router.Run(":8080")
}
