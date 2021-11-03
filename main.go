package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	username  string = "root"
	password  string = "password"
	addr      string = "127.0.0.1"
	port      int    = 3306
	database  string = "db01"
	parseTime bool   = false // time.time or string
)

type Product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Image    string `json:"image"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

func createTable(conn *sql.DB) error {
	sql := `CREATE TABlE table01(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(16) NOT NULL DEFAULT "",
		price INT DEFAULT 0,
		image VARCHAR(64) DEFAULT "",
		created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`
	_, err := conn.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

func insert(conn *sql.DB, p *Product) error {
	//"INSERT INTO table01(name, price) VALUES ('iphone12_64g', 26900)"
	_, err := conn.Exec("INSERT INTO table01(name, price, image) VALUES (?, ?, ?)", p.Name, p.Price, p.Image) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	return nil
}

func update(conn *sql.DB, p *Product) error {
	_, err := conn.Exec("UPDATE table01 SET name=?, price=?, image=? WHERE id=?", p.Name, p.Price, p.Image, p.ID) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	return nil
}

func delete(conn *sql.DB, id int) error {
	_, err := conn.Exec("DELETE FROM table01 WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func sqlSelect(conn *sql.DB, command string) ([]Product, error) {
	res, err := conn.Query(command)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var products []Product
	for res.Next() {
		var product Product
		err = res.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.CreateAt, &product.UpdateAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func main() {
	// connect to mariaDB
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v", username, password, addr, port, database, parseTime)
	conn, err := sql.Open("mysql", connInfo)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// http service
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	v1 := router.Group("/products")
	{
		v1.GET("/", func(ctx *gin.Context) {
			products, err := sqlSelect(conn, "SELECT * FROM table01")
			if err != nil {
				fmt.Println(err)
			}
			data, err := json.Marshal(products)
			if err != nil {
				fmt.Println(err)
			}
			ctx.Data(200, "application/json", data)
		})

		v1.POST("/", func(ctx *gin.Context) {
			var product Product
			ctx.BindJSON(&product)
			fmt.Println(product)

			if err := insert(conn, &product); err != nil {
				log.Println(err)
				ctx.String(500, "fail")
			}
			ctx.String(200, "ok")
		})

		v1.GET("/:productID", func(ctx *gin.Context) {
			productID := ctx.Param("productID")
			fmt.Println(productID)
			//ctx.File("./static/img/" + productID)
			ctx.String(200, "ok")
		})

		v1.DELETE("/:productID", func(ctx *gin.Context) {
			fmt.Println("delete id:", ctx.Param("productID"))
			productID, err := strconv.Atoi(ctx.Param("productID"))
			if err != nil {
				log.Println(err)
			}
			if err := delete(conn, productID); err != nil {
				log.Println(err)
			}
			ctx.String(200, "ok")
		})

		v1.PUT("/:productID", func(ctx *gin.Context) {
			var product Product
			ctx.BindJSON(&product)
			if err := update(conn, &product); err != nil {
				log.Println(err)
			}
			ctx.String(200, "ok")
		})

		v1.GET("/:productID/image", func(ctx *gin.Context) {
			productID := ctx.Param("productID")
			fmt.Println(productID, "image")
			//ctx.File("./static/img/" + productID)
			ctx.String(200, "ok")
		})
	}

	router.Static("/static", "./static")
	router.Run(":5000")
}
