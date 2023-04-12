package main

import (
	con "employee/pkg/config"
	c "employee/pkg/controller"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()

	file, err := os.Open("employee.csv")
	if err != nil {
		fmt.Print(err)
	}

	con.ReadData(file)
	log.Printf("main, execution time %s\n", time.Since(start))
	r := gin.Default()

	r.GET("/employee", c.GetEmployees)
	r.GET("/employee/:id", c.GetEmployeesByID)
	r.POST("/employee", c.AddEmployee)
	r.PUT("/employee/:id", c.UpdateEmployee)
	r.DELETE("/employee/:id", c.DeleteEmployee)

	r.Run("localhost:1050")
}
