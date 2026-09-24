package main

import (
	"log"
	"os"

	"github.com/MGT06/EventHub_Backend.git/internal/config"
	"github.com/MGT06/EventHub_Backend.git/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return
	}

	pdb := config.NewPsqlDb(os.Getenv("DBUSER"), os.Getenv("DBPASS"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))

	pool, err := pdb.Connect()
	if err != nil {
		log.Println("Cannot Connect to DB\nReason: ", err.Error())
		return
	}

	r := gin.Default()

	router.MainRouter(r, pool)

	r.Run("localhost:9000")
}
