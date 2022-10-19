package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	type Status struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	}

	data := []Status{}
	
	router.LoadHTMLFiles("template/index.html")
	router.GET("/", func(ctx *gin.Context) {

		weatherData := Status{
			Water: rand.Intn(100),
			Wind:  rand.Intn(100),
		}

		waterValue := weatherData.Water
		windValue := weatherData.Wind

		latestStatus := "AMAN"
		if windValue > 6 {
			if waterValue > 6 {
				latestStatus = "AMAN"
			} else if waterValue > 8 {
				latestStatus = "SIAGA"
			} else {
				latestStatus = "BAHAYA"
			}
		}

		data = append(data, weatherData)

		file, _ := json.MarshalIndent(data, "", "")
		_ = ioutil.WriteFile("weather.json", file, 0644)

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":         "Weather Report",
			"current_stats": latestStatus,
			"data":          data,
			"latest_index": len(data)-1,
			"index": len(data),
		})
	})
	router.Run()
}
