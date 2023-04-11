package main

import (
	con "employee/pkg/config"
	c "employee/pkg/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	file, err := os.Open("employee.csv")
	if err != nil {
		fmt.Print(err)
	}

	con.ReadData(file)
	r := gin.Default()

	r.GET("/employee", c.GetEmployees)
	r.GET("/employee/:id", c.GetEmployeesByID)
	r.POST("/employee", c.AddEmployee)
	r.PUT("/employee/:id", c.UpdateEmployee)
	r.DELETE("/employee/:id", c.DeleteEmployee)

	r.Run("localhost:7050")
}
