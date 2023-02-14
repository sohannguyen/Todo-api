package main

import (
	config2 "Todo-api/src/config"
	"Todo-api/src/config/database"
	"Todo-api/src/routers"
	"fmt"
	"log"
)

func init() {
	configDB, err := config2.LoadConfigDB(".")
	if err != nil {
		log.Fatalln("Can not load config: ", err)
	}
	fmt.Println(configDB)
	database.ConnectDB(&configDB)

	//migrate

}

func main() {
	r := routers.NewRoutes()
	r.Run()
}
