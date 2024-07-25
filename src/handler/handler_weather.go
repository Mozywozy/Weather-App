package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

const apiKey = "2f2cb1e909f04c137b5e16caa6620dbb"

type WeatherResponse struct {
    Name     string `json:"name"`
    Main     struct {
        Temp     float64 `json:"temp"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
        Icon        string `json:"icon"`
    } `json:"weather"`
    Wind struct {
        Speed float64 `json:"speed"`
    } `json:"wind"`
    Visibility int `json:"visibility"`
}

type ForecastResponse struct {
    List []struct {
        Dt   int64 `json:"dt"`
        Main struct {
            Temp float64 `json:"temp"`
        } `json:"main"`
        Weather []struct {
            Description string `json:"description"`
            Icon        string `json:"icon"`
        } `json:"weather"`
    } `json:"list"`
}

func GetWeather(c *gin.Context) {
    city := c.Query("city")
    if city == "" {
        city = "Jakarta"
    }

    weatherUrl := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey + "&units=metric"
    forecastUrl := "https://api.openweathermap.org/data/2.5/forecast?q=" + city + "&appid=" + apiKey + "&units=metric"

    weatherResp, err := http.Get(weatherUrl)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer weatherResp.Body.Close()

    if weatherResp.StatusCode != http.StatusOK {
        c.JSON(weatherResp.StatusCode, gin.H{"error": "Failed to get weather data"})
        return
    }

    forecastResp, err := http.Get(forecastUrl)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer forecastResp.Body.Close()

    if forecastResp.StatusCode != http.StatusOK {
        c.JSON(forecastResp.StatusCode, gin.H{"error": "Failed to get forecast data"})
        return
    }

    var weather WeatherResponse
    if err := json.NewDecoder(weatherResp.Body).Decode(&weather); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var forecast ForecastResponse
    if err := json.NewDecoder(forecastResp.Body).Decode(&forecast); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var weeklyForecast []gin.H
    for i, item := range forecast.List {
        if i%8 == 0 { // Get daily forecast (every 8th item corresponds to a new day)
            day := time.Unix(item.Dt, 0).Weekday().String()
            weeklyForecast = append(weeklyForecast, gin.H{
                "day":         day,
                "temp":        fmt.Sprintf("%.1f", item.Main.Temp), // Round to 1 decimal place
                "description": item.Weather[0].Description,
                "icon":        item.Weather[0].Icon,
            })
        }
    }

    c.HTML(http.StatusOK, "index.html", gin.H{
        "city":        weather.Name,
        "temp":        fmt.Sprintf("%.1f", weather.Main.Temp), // Round to 1 decimal place
        "description": weather.Weather[0].Description,
        "icon":        weather.Weather[0].Icon,
        "humidity":    weather.Main.Humidity,
        "windSpeed":   fmt.Sprintf("%.1f", weather.Wind.Speed), // Round to 1 decimal place
        "visibility":  weather.Visibility / 1000,
        "uvIndex":     "8", // Example value, replace with actual UV index data if available
        "airQuality":  "Good", // Example value, replace with actual air quality data if available
        "forecast":    weeklyForecast,
    })
}
