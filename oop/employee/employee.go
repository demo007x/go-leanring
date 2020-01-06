package employee

import "fmt"

type employee struct {
	firstName string
	lastName string
	totalLeavers int
	leaversTaken int
}

func (e *employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, e.totalLeavers - e.leaversTaken)
}

func New(firstName, lastName string, totalLeave, leavesTaken int) employee {
	e := employee{
		firstName:   firstName,
		lastName:     lastName,
		totalLeavers: totalLeave,
		leaversTaken: leavesTaken,
	}

	return e
}