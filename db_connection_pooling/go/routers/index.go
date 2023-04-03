package routers

import (
	"database/sql"
	"db_cnn_pooling/db"
	"db_cnn_pooling/models"
	"db_cnn_pooling/server"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

var query = "select * from products limit 1000"

var normalTime int64 = 0
var normalCount int64 = 0

var newTime int64 = 0
var newCount int64 = 0

var poolTime int64 = 0
var poolCount int64 = 0

func scanProduct(rows *sql.Rows) ([]*models.Products, error) {
	defer rows.Close()

	products := make([]*models.Products, 0)
	for rows.Next() {
		var p models.Products
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func GetProductNew() gin.HandlerFunc {
	dsn := db.NewPostgreSql()
	return func(c *gin.Context) {
		timeStart := time.Now()

		connNew, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Connect to postgres failed: New: ", err.Error())
		}

		rows, err := connNew.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		prods, err := scanProduct(rows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		elapsed := time.Since(timeStart).Microseconds()
		newCount++
		newTime += elapsed
		ave := float64(newTime / newCount)
		c.JSON(http.StatusOK, models.Response{
			Number:   newCount,
			Elapsed:  elapsed,
			Average:  ave,
			Products: prods,
		})
	}
}

func GetProductNormal(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStart := time.Now()

		rows, err := server.ConnNormal.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		prods, err := scanProduct(rows)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		elapsed := time.Since(timeStart).Microseconds()
		normalCount++
		normalTime += elapsed
		ave := float64(normalTime / normalCount)
		c.JSON(http.StatusOK, models.Response{
			Number:   normalCount,
			Elapsed:  elapsed,
			Average:  ave,
			Products: prods,
		})
	}
}

func GetProductPooled(server server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStart := time.Now()
		rows, err := server.ConnPooled.Query(query)

		if err != nil {
			log.Println("err server.ConnPooled: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		prods, err := scanProduct(rows)
		if err != nil {
			log.Println("err scanProduct: ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		elapsed := time.Since(timeStart).Microseconds()
		poolCount++
		poolTime += elapsed
		ave := float64(poolTime / poolCount)
		c.JSON(http.StatusOK, models.Response{
			Number:   poolCount,
			Elapsed:  elapsed,
			Average:  ave,
			Products: prods,
		})
	}
}
