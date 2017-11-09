package main

import "github.com/gin-gonic/gin"

func initRouter()  {
	router := gin.Default()

	// act browser http
	router.Group("/api/browser/act")
	{
		//actBroswer.POST("/login", loginEndpoint)
	}

	//act wallet http
	router.Group("/api/wallet/act")
	{

	}
}
