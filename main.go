package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"practive_service/config/db"
	"practive_service/context"
	"practive_service/infrastucture/models"
	"practive_service/routers"
	"practive_service/utils"
)

func main() {
	postgresUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		utils.GetEnv("PG_HOST", "localhost"),
		utils.GetEnv("PG_USER", "postgres"),
		utils.GetEnv("PG_PASSWORD", "123456"),
		utils.GetEnv("PG_DBNAME", "practice"),
		utils.GetEnv("PG_PORT", "5432"))

	dbConn := db.Connect(postgresUrl)
	dbConn.AutoMigrate(&models.Person{})
	defer dbConn.Close()

	app, err := context.BuildApplication(dbConn)
	if err != nil {
		log.Fatalln(os.Stderr, err)
	}

	r := routers.SetupRouters(app)
	r.Run(":8080")
}
