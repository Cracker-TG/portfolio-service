package server

import (
	"fmt"
	"log"

	"github.com/Cracker-TG/portfolio-service/config"
	"github.com/Cracker-TG/portfolio-service/database"
	"github.com/Cracker-TG/portfolio-service/routes"
)

var DBinstance database.IDBinstance

func Init() {
	r := routes.NewRouter()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Print(config)

	DBinstance = &database.DBinstance{}
	DBinstance.InitDB(config.MONGO_HOST, config.MONGO_PORT, config.MOGO_DB)
	r.Run(":" + config.PORT)
}
