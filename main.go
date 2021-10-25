package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func create(conn *sql.DB) error {
	sql := `CREATE TABlE table01(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(16) NOT NULL,
		price INT,
		image VARCHAR(64),
		PRIMARY KEY (id)
	);`
	_, err := conn.Query(sql)
	if err != nil {
		return err
	}
	return nil
}

func insert(conn *sql.DB, command string) error {
	_, err := conn.Query(command) //("INSERT INTO user_info (name, age) VALUES (?, ?)","syhlion",18,)
	if err != nil {
		return err
	}
	//defer res.Close()
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
		var (
			id    int
			name  string
			price sql.NullInt64
			image sql.NullString
		)
		err = res.Scan(&id, &name, &price, &image)
		if err != nil {
			return err
		}
		var product = Product{id, name, int(price.Int64), image.String}
		products = append(products, product)
		fmt.Println(product.Name, product.Price)
	}
	fmt.Println(products)
	return nil
}

func main() {
	// connect to mariaDB
	conn, err := sql.Open("mysql", "root:951219@tcp(localhost:3306)/db01")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	err = create(conn)

	// err = insert(conn, "INSERT INTO table01(name, price) VALUES ('iphone12_64g', 26900)")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err = sqlSelect(conn, "SELECT * FROM table01")
	if err != nil {
		fmt.Println(err)
	}

}
