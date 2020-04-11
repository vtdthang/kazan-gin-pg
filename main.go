package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/vtdthang/kazan-gin-pg/product"
	router "github.com/vtdthang/kazan-gin-pg/routers"
)

func initDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := InitProductAPI(db)

	r := gin.Default()

	router.PrepareProductRoutes(r, productAPI)
	router.PrepareUserRoutes(r)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
