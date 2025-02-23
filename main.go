package main

import (
	"challenge2016/routes"
	"challenge2016/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	_, err = services.LoadCities("cities.csv")
	if err != nil {
		fmt.Println("Error loading cities:", err)
		return
	}
	fmt.Println("Cities loaded successfully.")

	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.GET("/check", routes.CheckDistribution)
	router.POST("/distributor", routes.AddDistributor)
	router.POST("/permission", routes.AddPermission)

	// please refer to the documentation :
	// this is the api documentation link for better understanding of the API routes: "https://documenter.getpostman.com/view/25819639/2sAYdcrCTq"
	router.Run(":8080")
}
