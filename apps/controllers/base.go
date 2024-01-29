package controllers

import (
	"TokoOnline/apps/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func (server *Server) Initalize(database DatabaseConfig) {
	server.Router = mux.NewRouter()
	server.InitalizeRouters()
	fmt.Println("Connect To Database")
	var err error
	dsn := database.Username + ":" + database.Password + "@tcp(" + database.Host + ":" + strconv.Itoa(database.Port) + ")/" + database.DBName + "?parseTime=true"
	server.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//Auto Imigrate database
	server.Db.AutoMigrate(&models.User{})
	server.Db.AutoMigrate(&models.Address{})
	server.Db.AutoMigrate(&models.Category{})
	// server.Db.AutoMigrate(&models.Order{})
	// server.Db.AutoMigrate(&models.OrderItem{})
	// server.Db.AutoMigrate(&models.OrderCustomer{})
	server.Db.AutoMigrate(&models.Product{})
	// server.Db.AutoMigrate(&models.Payment{})
	server.Db.AutoMigrate(&models.ProductImage{})
	// server.Db.AutoMigrate(&models.Shipment{})
	server.Db.AutoMigrate(&models.Section{})
	// server.Db.AutoMigrate(&models.Cart{})
	// server.Db.AutoMigrate(&models.CartItem{})
	// server.Db.AutoMigrate(&models.OrderCustomer{})
	// server.Db.AutoMigrate(&models.{})
	// server.Db.AutoMigrate(&models.{})

	// seeders.DbSeeder(server.Db)

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))

}
