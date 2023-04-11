package controller

import (
	con "employee/pkg/config"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID    int
	Name  string `json:"Name"`
	Dept  string `json:"Dept"`
	Title string `json:"Title"`
	Comp  string `json:"Comp"`
}

var db = con.CreateConn()

func GetEmployees(c *gin.Context) {
	query := "SELECT * FROM employee"

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Dept, &employee.Title, &employee.Comp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		employees = append(employees, employee)
	}

	c.JSON(http.StatusOK, employees)
}

func GetEmployeesByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rows, _ := db.Query("SELECT * FROM employee where ID = ?", id)
	var employee Employee
	for rows.Next() {
		// var ID int
		// var Name string
		// var Dept string
		// var Title string
		// var Comp string

		rows.Scan(&employee.ID, &employee.Name, &employee.Dept, &employee.Title, &employee.Comp)

		if employee.Name != "" {
			c.IndentedJSON(http.StatusOK, gin.H{"1. ID": id, "2. Name": employee.Name, "3. Dept": employee.Dept, "4. Title": employee.Title, "5. Comp": employee.Comp})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

func AddEmployee(c *gin.Context) {
	// 	var employee Employee
	// 	err := c.BindJSON(&employee)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	result, err := db.Exec("INSERT INTO employee (ID, Name, Dept, Title, Comp) VALUES (?, ?, ?, ?, ?)", employee.ID, employee.Name, employee.Dept, employee.Title, employee.Comp)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	id, err := result.LastInsertId()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	employee.ID = string(rune(id))
	// 	c.JSON(http.StatusCreated, employee)
	// }

	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO employee (ID, Name, Dept, Title, Comp) VALUES (?, ?, ?, ?, ?)", employee.ID, employee.Name, employee.Dept, employee.Title, employee.Comp)
	if err != nil {
		fmt.Print(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	employee.ID = int(id)
	c.JSON(http.StatusCreated, employee)
}

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var employee Employee
	err = c.BindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE employee SET ID = ?,  Name = ?, Dept = ?, Title = ?, Comp = ? WHERE ID = ?"

	result, err := db.Exec(query, employee.ID, employee.Name, employee.Dept, employee.Title, employee.Comp, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}

	employee.ID = id
	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	result, err := db.Exec("DELETE FROM employee WHERE ID = ?", id)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(result)
}
