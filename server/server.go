package server

import (
	"os"

	"github.com/Cracker-TG/portfolio-service/routes"
)

func Init() {
	r := routes.NewRouter()
	r.Run(":" + os.Getenv("GOPORT"))
}
