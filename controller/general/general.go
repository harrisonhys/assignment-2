package general

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Orderan struct {
	ID            int
	Customer_name string
	Ordered_at    string
}

type OrderDetail struct {
	ID          int
	Item_code   string
	Description string
	Quantity    int
	Order_id    int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "orders_by"
)

var (
	db  *sql.DB
	err error
)

func conn() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
}

func close() {
	defer db.Close()
}

func createData(c *gin.Context) {

	conn()

	customerName := c.PostForm("customerName")
	orderedAt := c.PostForm("orderedAt")

	fmt.Println(jsonData)

	close()

}

func GetData(c *gin.Context) {

	conn()

	c.Header("Context-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")

	var results = []Orderan{}

	sqlStatement := `SELECT * from orders`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var Orderan = Orderan{}

		err = rows.Scan(&Orderan.ID, &Orderan.Customer_name, &Orderan.Ordered_at)

		if err != nil {
			panic(err)
		}

		results = append(results, Orderan)
	}

	close()

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": results,
	})

}

func ShowData(c *gin.Context) {

	conn()

	var results = []Orderan{}
	var detail = []OrderDetail{}

	sqlStatement := `SELECT * from orders WHERE order_id= $1`
	sqlStatementDetail := `SELECT * from items WHERE order_id= $1`

	ID := c.Param("id")

	rows, err := db.Query(sqlStatement, ID)
	rowsDetail, errDetail := db.Query(sqlStatementDetail, ID)

	if err != nil {
		panic(err)
	}

	if errDetail != nil {
		panic(errDetail)
	}

	defer rows.Close()

	for rows.Next() {
		var Orderan = Orderan{}

		err = rows.Scan(&Orderan.ID, &Orderan.Customer_name, &Orderan.Ordered_at)

		if err != nil {
			panic(err)
		}

		results = append(results, Orderan)
	}

	close()

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": results,
	})

}

func updateData() {

	sqlStatement := `UPDATE Orderans SET full_name = $2, email = $3, age = $4, division= $5 WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1, "usertest", "usertest@gmail.com", 23, "IT")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Update Data Amount : ", count)
}

func deleteData() {

	sqlStatement := `DELETE from Orderans WHERE id = $1;`

	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Delete Data Amount : ", count)

}
