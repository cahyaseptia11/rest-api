package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//car

// {
// 	"id" : "1",
// 	"brand" : "Honda",
// 	"type" : "city"

// }

type car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"car_type"`
}

var cars = []car{
	{ID: "1", Brand: "Honda", Type: "City"},
	{ID: "2", Brand: "Toyota", Type: "Avanza"},
	{ID: "3", Brand: "Lifan", Type: "SUV"},
}

func main() {
	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	//GET / cars - list cars
	r.GET("/cars", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, cars)
	})

	//POST / cars - create car
	r.POST("/cars", func(ctx *gin.Context) {
		var car car
		if err := ctx.ShouldBindJSON(&car); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		cars = append(cars, car)
		ctx.JSON(http.StatusCreated, car)
	})

	//DELETE / cars/:id - delete car
	r.DELETE("/cars/:car_id", func(ctx *gin.Context) {
		id := ctx.Param("car_id")
		for i, car := range cars {
			if car.ID == id {
				cars = append(cars[:i], cars[i+1:]...)
				break
			}
		}
		ctx.Status(http.StatusNoContent)
	})

	//GET / cars/:id - get one car
	r.GET("cars/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, a := range cars {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	//PUT / cars/:id - update car
	r.PUT("/cars/:car_id", func(c *gin.Context) {
		var newCar car
		id := c.Param("id")
		if id == "" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Insert ID"})
		} else {
			c.BindJSON(&newCar)
			// return
		}
		cars = append(cars, newCar)
		c.IndentedJSON(http.StatusCreated, newCar)
	})

	r.Run()

}
