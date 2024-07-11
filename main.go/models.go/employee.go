package models

import (
	"github.com/google/uuid"
)

type Employee struct {
	EmpId                  uuid.UUID `json:"emp_id"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	Gender                 string    `json:"gender"`
	PhoneNumber            string    `json:"phone_number"`
	EmployeeEmail          string    `json:"employee_email"`
	Address                string    `json:"address"`
	BloodGroup             string    `json:"blood_group"`
	EmergencyContactNumber string    `json:"emergency_contact_number"`
}
