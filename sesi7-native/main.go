package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println()

}

const (
	host     = "localhost"
	port     = 5432
	user     = "esra"
	password = "koinworks"
	db_name  = "hacktiv_sesi7"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	//create employee
	// emp := Employee{
	// 	Email:    "fadia@gmail.com",
	// 	FullName: "Fadia",
	// 	Age:      21,
	// 	Division: "Accounting Manager",
	// }

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	//update employee
	emp := Employee{
		ID:       1,
		FullName: "Melati Daeva",
		Email:    "melati@gmail.com",
		Age:      27,
		Division: "Olahraga",
	}
	fmt.Println("====== Update Employee by id 1 ======")
	err = updateEmployee(db, &emp)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	fmt.Println("update data success", emp)

	employees, err := getAllEmployees(db)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for _, employee := range *employees {
		employee.Print()
	}
	fmt.Println("====== Delete Employee by id 6 ======")
	isSucces, err := deleteEmployee(db, 6)
	if err != nil {
		fmt.Println("Error :", err.Error())
		return
	}
	fmt.Println(isSucces)

	fmt.Println("====== Get Employee by id 4 ======")
	employee, err := getEmployeeById(db, 4)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// res,err := deleteEmployee(db,&res)
	// if err != nil {
	// 	fmt.Print("error: ", err.Error())
	// 	return
	// }

	employee.Print()
}

func connectDB() (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID, &employee.FullName,
			&employee.Email, &employee.Age, &employee.Division,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil

}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES($1, $2, $3, $4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()

}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func updateEmployee(db *sql.DB, request *Employee) error {
	query := `
		UPDATE employees
		SET full_name=$2, email=$3,division=$4,age=$5
		WHERE id=$1;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(request.ID, request.FullName, request.Email, request.Division, request.Age)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()

}

func deleteEmployee(db *sql.DB, id int8) (bool, error) {
	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_ = stmt.QueryRow(id)

	return true, nil
}
