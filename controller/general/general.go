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

type Order struct {
	Orderedat string
	Custname  string
	Items     []Items
}

type Items struct {
	Itemcode    string
	Description string
	Quantity    int
}

func CreateData(c *gin.Context) {

	conn()

	var requestBody Order

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("Exec err:", err)
	}

	var ordered_at = requestBody.Orderedat
	var customer_name = requestBody.Custname
	var items = requestBody.Items

	order_id := 0
	err := db.QueryRow("INSERT INTO orders (customer_name,ordered_at) VALUES ($1,$2) Returning order_id", customer_name, ordered_at).Scan(&order_id)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(items); i++ {
		var item_code = items[i].Itemcode
		var item_desc = items[i].Description
		var qty = items[i].Quantity

		result, erro := db.Exec("INSERT INTO items (item_code,description,quantity,order_id) VALUES ($1,$2,$3,$4)", item_code, item_desc, qty, order_id)
		if erro != nil {
			fmt.Println("Exec err:", erro.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": result,
		})

	}

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

	// conn()

	// var results = []Orderan{}
	// var detail 	= []OrderDetail{}

	// sqlStatement := `SELECT * from orders WHERE order_id= $1`
	// sqlStatementDetail := `SELECT * from items WHERE order_id= $1`

	// ID := c.Param("id")

	// rows, err := db.Query(sqlStatement, ID)
	// rowsDetail, errDetail := db.Query(sqlStatementDetail, ID)

	// if err != nil {
	// 	panic(err)
	// }

	// if errDetail != nil {
	// 	panic(errDetail)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var Orderan = Orderan{}

	// 	err = rows.Scan(&Orderan.ID, &Orderan.Customer_name, &Orderan.Ordered_at)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	results = append(results, Orderan)
	// }

	// close()

	// c.JSON(http.StatusOK, gin.H{
	// 	"status":   http.StatusOK,
	// 	"response": results,
	// })

}

func updateData(c *gin.Context) {

	conn()

	id := c.Param("id")

	var requestBody Order

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("Exec err:", err)
	}

	var ordered_at = requestBody.Orderedat
	var customer_name = requestBody.Custname
	var items = requestBody.Items

	err := db.QueryRow("UPDATE orders set customer_name = $1, ordered_at = $2 WHERE order_id = $3", customer_name, ordered_at, id)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(items); i++ {
		var item_code = items[i].Itemcode
		var item_desc = items[i].Description
		var qty = items[i].Quantity

		result, erro := db.Exec("INSERT INTO items (item_code,description,quantity,order_id) VALUES ($1,$2,$3,$4)", item_code, item_desc, qty, order_id)
		if erro != nil {
			fmt.Println("Exec err:", erro.Error())
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"response": result,
		})

	}

	close()
}

func DeleteData(c *gin.Context) {

	id := c.Param("id")
	sqlStatement1 := `DELETE from orders WHERE order_id = $1;`
	sqlStatement2 := `DELETE from items WHERE order_id = $1;`

	res1, err1 := db.Exec(sqlStatement1, id)
	res2, err2 := db.Exec(sqlStatement2, id)

	if err1 != nil {
		panic(err1)
	}

	if err2 != nil {
		panic(err2)
	}

	fmt.Println(res1)
	fmt.Println(res2)

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Berhasil Delete Data",
	})

}
