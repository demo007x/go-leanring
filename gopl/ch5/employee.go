package ch5

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerID int
}

var dilbert Employee
var Point struct{x, y int}

func main() {
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	id := dilbert.ID
	EmployeeByID(id).Salary = 0


}

func EmployeeByID(id int) *Employee {

}
