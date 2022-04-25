package main

import (
	"api/handler"
	"api/penyakit"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){

	dsn := "root:@tcp(localhost:3306)/dnamatching?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(penyakit.Penyakit{})

	// API layering
	// Request -> Handler -> Service -> Repository -> Database
	penyakitRepository := penyakit.NewRepository(db)
	penyakitService := penyakit.NewService(penyakitRepository)
	penyakitHandler := handler.NewPenyakitHandler(penyakitService)

	router := gin.Default()

	v1 := router.Group("/v1")

	// v1.GET("/penyakit", penyakitHandler.GetAllPenyakitHandler)
	// v1.GET("/penyakit/:id", penyakitHandler.GetPenyakitHandler)

	v1.POST("/penyakit", penyakitHandler.CreatePenyakitHandler)
	// v1.PUT("/penyakit/:id", penyakitHandler.UpdatePenyakitHandler)
	// v1.DELETE("/penyakit/:id", penyakitHandler.DeletePenyakitHandler)
	
	router.Run()
}



