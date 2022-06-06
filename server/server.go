package server

import (
	"log"

	"github.com/Cracker-TG/crboard/config"
	"github.com/Cracker-TG/crboard/database"
	"github.com/Cracker-TG/crboard/routes"
)

var DBinstance database.IDBinstance

func Init() {
	r := routes.NewRouter()
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DBinstance = &database.DBinstance{}
	DBinstance.InitDB(config.MONGO_HOST, config.MONGO_PORT, config.MOGO_DB)
	r.Run(":" + "3000")
}
