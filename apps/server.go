package apps

import (
	"TokoOnline/apps/controllers"
	"fmt"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
}

func Run() {
	var server = controllers.Server{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Err Load env File")
	}
	AppPort := os.Getenv("port")

	dbConfig := controllers.DatabaseConfig{
		Host: os.Getenv("DBHOST"),
	}
	portStr := os.Getenv("DBPORT")
	dbConfig.Port, _ = strconv.Atoi(portStr)
	dbConfig.Username = os.Getenv("DBUSER")
	dbConfig.Password = os.Getenv("DBPASS")
	dbConfig.DBName = os.Getenv("DBNAME")

	server.Initalize(dbConfig)
	server.Run(":" + AppPort)

}
