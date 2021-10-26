package main

import (
	"database/sql"
	"fmt"

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
	Price    int    `json:"age"`
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

func insert(conn *sql.DB, name string, price int, image string) error {
	//"INSERT INTO table01(name, price) VALUES ('iphone12_64g', 26900)"
	_, err := conn.Exec("INSERT INTO table01(name, price, image) VALUES (?, ?, ?)", name, price, image) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	return nil
}

func update(conn *sql.DB, id int, name string, price int, image string) error {
	_, err := conn.Exec("UPDATE table01 SET name=?, price=?, image=? WHERE id=?", name, price, image, id) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
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

func sqlSelect(conn *sql.DB, command string) error {
	res, err := conn.Query(command)
	if err != nil {
		return err
	}
	defer res.Close()

	var products []Product
	for res.Next() {
		var product Product
		err = res.Scan(&product.ID, &product.Name, &product.Price, &product.Image, &product.CreateAt, &product.UpdateAt)
		if err != nil {
			return err
		}
		products = append(products, product)
		//fmt.Println(product.Name, product.Price)
	}
	fmt.Println(products)
	return nil
}

func main() {
	// connect to mariaDB
	connInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=%v", username, password, addr, port, database, parseTime)
	conn, err := sql.Open("mysql", connInfo)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// err = createTable(conn)
	// if err != nil {
	// 	log.Println(err)
	// }

	// err = insert(conn, "iphone11", 30000, "")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = update(conn, 2, "iphone11", 3000, "")
	// if err != nil{
	// 	log.Println(err)
	// }

	// err = delete(conn, 1)
	// if err != nil {
	// 	log.Println(err)
	// }

	err = sqlSelect(conn, "SELECT * FROM table01")
	if err != nil {
		fmt.Println(err)
	}

}
