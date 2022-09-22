package general

import(
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "password"
	dbname		= "db-go-sql"
)

type Employee struct {
	ID				int
	Full_name		string
	Email			string
	Age				int
	Division		string
}

var (
	db *sql.DB
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

func createData() {
	conn()

	var employee = Employee{}

	sqlStatement := `INSERT INTO employees (full_name,email,age,division)
					VALUES ($1,$2,$3,$4) 
					Returning *
					`

	err = db.QueryRow(sqlStatement, "usertest1", "usertest1@gmail.com", 24, "IT").
	Scan(&employee.ID,&employee.Full_name,&employee.Email,&employee.Age,&employee.Division)

	if err != nil {
		panic(err)
	}

	close()

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetData(c *gin.Context) {

	conn()

	c.Header("Context-Type", "application/x-www-form-urlencoded")
	c.Header("Access-Control-Allow-Origin", "*")
	
	var results = []Employee{}
	sqlStatement := `SELECT * from employees`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID,&employee.Full_name,&employee.Email,&employee.Age,&employee.Division)
		
		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	close()

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": results,
	})

}

func showData(ID int) {

	var results = []Employee{}
	sqlStatement := `SELECT * from employees WHERE id=$1`

	rows, err := db.Query(sqlStatement, ID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID,&employee.Full_name,&employee.Email,&employee.Age,&employee.Division)
		
		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)

}

func updateData() {

	sqlStatement := `UPDATE employees SET full_name = $2, email = $3, age = $4, division= $5 WHERE id = $1;`

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

	sqlStatement := `DELETE from employees WHERE id = $1;`

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