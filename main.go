package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var server Server = Server{router: gin.Default()}

	data, err := GetJson("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Println("Failed to pull url, exiting")
		os.Exit(1)
	}

	astronauts := AstroList{}
	err = json.Unmarshal(data, &astronauts)
	if err != nil {
		fmt.Println(err)
		return
	}
	ships := TransformAstroData(astronauts)
	server.ships = ships

	fmt.Println(ships)
	server.initializeRoutes()

	server.router.Run()
}
