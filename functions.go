package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJson(url string) ([]byte, error) {
	data, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()
	body, err := io.ReadAll(data.Body)
	return body, err

}

func TransformAstroData(people AstroList) *[]ShipCrew {
	shipMap := make(map[string][]string)
	var tmp []ShipCrew

	for i := range people.People {
		shipMap[people.People[i].Craft] = append(shipMap[people.People[i].Craft], people.People[i].Name)
		fmt.Println(people.People[i])
	}

	for k, v := range shipMap {
		var shipCrew ShipCrew
		shipCrew.Name = k
		shipCrew.Crew = v
		tmp = append(tmp, shipCrew)
	}

	return &tmp
}

func (s *Server) ShowMeTheCrew(c *gin.Context) {
	c.JSON(200, s.ships)
}

func InterestingResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "interdasting",
	})
}

func LikingThisSoFar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Don't forget to feed the cat!!!",
	})
}
