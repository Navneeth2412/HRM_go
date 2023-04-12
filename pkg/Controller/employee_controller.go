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

	rows, err := db.Query("SELECT * FROM employee")
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

	c.IndentedJSON(http.StatusOK, employees)
}

func GetEmployeesByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rows, _ := db.Query("SELECT * FROM employee where ID = ?", id)
	var employee Employee
	for rows.Next() {

		rows.Scan(&employee.ID, &employee.Name, &employee.Dept, &employee.Title, &employee.Comp)

		if employee.Name != "" {
			c.IndentedJSON(http.StatusOK, gin.H{"1. ID": id, "2. Name": employee.Name, "3. Dept": employee.Dept, "4. Title": employee.Title, "5. Comp": employee.Comp})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

func AddEmployee(c *gin.Context) {

	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO employee (ID, Name, Dept, Title, Comp) VALUES (?, ?, ?, ?, ?)", employee.ID, employee.Name, employee.Dept, employee.Title, employee.Comp)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(result)
	add := "ADDED!!!"
	c.JSON(http.StatusCreated, add)

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

	result, err := db.Exec("UPDATE employee SET ID = ?,  Name = ?, Dept = ?, Title = ?, Comp = ? WHERE ID = ?", employee.ID, employee.Name, employee.Dept, employee.Title, employee.Comp, id)
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

	update := "UPDATED!!!"
	c.JSON(http.StatusCreated, update)

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
	delete := "DELETED!!!"
	c.JSON(http.StatusOK, delete)
}
