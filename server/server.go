package server

import (
	"log"

	"github.com/Cracker-TG/portfolio-service/config"
	"github.com/Cracker-TG/portfolio-service/database"
	"github.com/Cracker-TG/portfolio-service/middlewares"
	"github.com/Cracker-TG/portfolio-service/routes"
	"github.com/gin-gonic/gin"
)

var DBinstance database.IDBinstance

func Init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	router := gin.New()
	router.Use(middlewares.Cors(&config))
	r := routes.NewRouter(router)

	DBinstance = &database.DBinstance{}
	DBinstance.InitDB(config.MONGO_HOST, config.MONGO_PORT, config.MOGO_DB)

	r.Run(":" + config.PORT)
}
