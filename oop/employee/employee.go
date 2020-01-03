package employee

import "fmt"

type Employee struct {
	FirstName string
	LastName string
	TotalLeavers int
	LeaversTaken int
}

func (e *Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, e.TotalLeavers - e.LeaversTaken)
}

func main() {

}