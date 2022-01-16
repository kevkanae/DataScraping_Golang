package main

import (
	"fmt"
	"lead/routes"
	"lead/utils"
)

func main() {

	//Gin Server
	server := routes.SetupRouter()
	err := server.Run()
	if err != nil {
		fmt.Println(utils.Wrap(err, "Couldn't Start Gin Server"))
		return
	}
}
