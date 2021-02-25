package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type service struct {
	gorm.Model

	Id      string `json:id`
	Title   string `json:title`
	Cost    string `json:cost`
	Calibre string `json:calibre`
}

var db *gorm.DB
var err error

func main() {
	router := gin.Default()

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=Organization sslmode=disable password=Verma_123")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()

	db.AutoMigrate(&service{})

	router.POST("/add", func(c *gin.Context) {
		var p service
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"title":   p.Title,
				"cost":    p.Cost,
				"calibre": p.Calibre,
			})
			db := dbConn()
			insForm, err := db.Prepare("INSERT INTO service(title, cost, calibre) VALUES(?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(p.Title, p.Cost, p.Calibre)
			fmt.Printf("title: %s; cost: %s; calibre: %s", p.Title, p.Cost, p.Calibre)
		}

	})

	router.PUT("/updatetitle", func(c *gin.Context) {

		var p service
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"title": p.Title,
				//"cost":   p.Cost,
				//"calibre": p.Calibre,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE service SET title=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Title, p.Id)
			//fmt.Printf("title: %s; cost: %s; calibre: %s",  p.Cost, p.Calibre)
		}
	})

	router.PUT("/updatecost", func(c *gin.Context) {

		var p service
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"cost": p.Cost,
				//"cost":   p.Cost,
				//"calibre": p.Calibre,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE service SET cost=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Cost, p.Id)
			//fmt.Printf("title: %s; cost: %s; calibre: %s", p.Title, p.Cost, p.Calibre)
		}
	})

	router.PUT("/updatecalibre", func(c *gin.Context) {

		var p service
		if c.BindJSON(&p) == nil {

			c.JSON(200, gin.H{
				"calibre": p.Calibre,
				//"cost":   p.Cost,
				//"calibre": p.Calibre,
			})
			db := dbConn()
			upForm, err := db.Prepare("UPDATE service SET calibre=? Where id=?")
			if err != nil {
				panic(err.Error())
			}
			upForm.Exec(p.Calibre, p.Id)
			//fmt.Printf("title: %s; cost: %s; calibre: %s", p.Title, p.Cost, p.Calibre)
		}
	})

	router.GET("/GET", func(c *gin.Context) {
		var p service
		if c.BindJSON(&p) == nil {
			db := dbConn()
			selDB, err := db.Query("SELECT * FROM service WHERE id=?", p.Id)
			if err != nil {
				panic(err.Error())
			}

			var id, title, cost, calibre string
			for selDB.Next() {

				err = selDB.Scan(&id, &title, &cost, &calibre)
				if err != nil {
					panic(err.Error())
				}
			}
			fmt.Printf("title: %s; cost: %s; calibre: %s", title, cost, calibre)

			c.JSON(200, gin.H{
				"id":      id,
				"title":   title,
				"cost":    cost,
				"calibre": calibre,
			})

		}

	})

	router.GET("/getall", func(c *gin.Context) {
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM service")
		if err != nil {
			panic(err.Error())
		}

		var id, title, cost, calibre string
		for selDB.Next() {

			err = selDB.Scan(&id, &title, &cost, &calibre)
			c.JSON(200, gin.H{
				"id":      id,
				"title":   title,
				"cost":    cost,
				"calibre": calibre,
			})
			fmt.Printf("title: %s; cost: %s; calibre: %s", title, cost, calibre)
			if err != nil {
				panic(err.Error())
			}
		}
		//fmt.Printf("title: %s; cost: %s; calibre: %s", title, cost, calibre)

	})

	router.DELETE("/delete", func(c *gin.Context) {
		var p service
		if c.BindJSON(&p) == nil {
			db := dbConn()
			delForm, err := db.Prepare("DELETE FROM service WHERE title=?")
			if err != nil {
				panic(err.Error())
			}
			delForm.Exec(p.Title)
			log.Println("DELETE")
			defer db.Close()
		}

	})

	router.Run(":8080")
}
