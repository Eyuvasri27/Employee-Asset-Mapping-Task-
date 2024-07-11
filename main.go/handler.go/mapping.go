package handlers

import (
	"employee-asset-mapping/db"
	"employee-asset-mapping/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AssignAssetMapping(w http.ResponseWriter, r *http.Request) {
	var mapping models.EmployeeAssetMapping
	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mapping.ID = uuid.New()
	_, err := db.DB.Exec(`INSERT INTO employee_asset_mapping (id, emp_id, asset_id) VALUES ($1, $2, $3)`, mapping.ID, mapping.EmpId, mapping.AssetId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapping)
}

func GetAllAssetsMappedToEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId, err := uuid.Parse(vars["employeeId"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	rows, err := db.DB.Query(`SELECT id, emp_id, asset_id FROM employee_asset_mapping WHERE emp_id=$1`, empId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mappings []models.EmployeeAssetMapping
	for rows.Next() {
		var mapping models.EmployeeAssetMapping
		err := rows.Scan(&mapping.ID, &mapping.EmpId, &mapping.AssetId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mappings = append(mappings, mapping)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mappings)
}

func RemoveAssetMapping(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mappingId, err := uuid.Parse(vars["mappingId"])
	if err != nil {
		http.Error(w, "Invalid mapping ID", http.StatusBadRequest)
		return
	}
	_, err = db.DB.Exec(`DELETE FROM employee_asset_mapping WHERE id=$1`, mappingId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT e.emp_id, e.first_name, e.last_name, e.gender, e.phone_number, e.employee_email, e.address, e.blood_group, e.emergency_contact_number, COUNT(m.asset_id) as asset_count
        FROM employees e
        LEFT JOIN employee_asset_mapping m ON e.emp_id = m.emp_id
        GROUP BY e.emp_id`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type DashboardEmployee struct {
		EmpId                  uuid.UUID `json:"emp_id"`
		FirstName              string    `json:"first_name"`
		LastName               string    `json:"last_name"`
		Gender                 string    `json:"gender"`
		PhoneNumber            string    `json:"phone_number"`
		EmployeeEmail          string    `json:"employee_email"`
		Address                string    `json:"address"`
		BloodGroup             string    `json:"blood_group"`
		EmergencyContactNumber string    `json:"emergency_contact_number"`
		AssetCount             int       `json:"asset_count"`
	}

	var employees []DashboardEmployee
	for rows.Next() {
		var emp DashboardEmployee
		err := rows.Scan(&emp.EmpId, &emp.FirstName, &emp.LastName, &emp.Gender, &emp.PhoneNumber, &emp.EmployeeEmail, &emp.Address, &emp.BloodGroup, &emp.EmergencyContactNumber, &emp.AssetCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"EmployeeList": employees,
	})
}
