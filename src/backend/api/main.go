package main

import (
	"api/handler"
	"api/penyakit"
	"api/riwayat"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "4isuGYYya5:Xxmdskrzr6@tcp(remotemysql.com:3306)/4isuGYYya5?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:@tcp(localhost:3306)/dnamatching?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(penyakit.Penyakit{})
	db.AutoMigrate(riwayat.Riwayat{})

	// API layering
	// Request -> Handler -> Service -> Repository -> Database
	penyakitRepository := penyakit.NewRepository(db)
	penyakitService := penyakit.NewService(penyakitRepository)
	penyakitHandler := handler.NewPenyakitHandler(penyakitService)

	riwayatRepository := riwayat.NewRepository(db)
	riwayatService := riwayat.NewService(riwayatRepository)
	riwayatHandler := handler.NewRiwayatHandler(riwayatService, penyakitService)

	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/v1")

	v1.POST("/penyakit", penyakitHandler.CreatePenyakitHandler)
	v1.POST("/riwayat", riwayatHandler.CreateRiwayatHandler)
	v1.GET("/riwayat", riwayatHandler.QueryRiwayatHandler)

	router.Run()
}
