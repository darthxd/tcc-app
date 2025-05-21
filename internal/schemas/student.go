package schemas

import "time"

type Student struct {
	ID        uint
	Name      string
	RM        string
	RA        string
	CPF       string
	Phone     string
	Birthdate time.Time
	Course    string
	Grade     string
	Email     string
	Password  string
	Photo     *string
	Biometry  *string
	InSchool  bool
}
