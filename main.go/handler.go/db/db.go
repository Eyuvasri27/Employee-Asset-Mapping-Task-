package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	employeeTable := `CREATE TABLE IF NOT EXISTS employees (
        emp_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        first_name VARCHAR(50) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        gender CHAR(1),
        phone_number VARCHAR(15),
        employee_email VARCHAR(100) UNIQUE NOT NULL,
        address TEXT,
        blood_group VARCHAR(3),
        emergency_contact_number VARCHAR(15)
    );`

	assetTable := `CREATE TABLE IF NOT EXISTS assets (
        asset_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        asset_name VARCHAR(100) NOT NULL,
        asset_type VARCHAR(50) NOT NULL
    );`

	mappingTable := `CREATE TABLE IF NOT EXISTS employee_asset_mapping (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        emp_id UUID REFERENCES employees(emp_id) ON DELETE CASCADE,
        asset_id UUID REFERENCES assets(asset_id) ON DELETE CASCADE
    );`

	_, err := DB.Exec(employeeTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(assetTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(mappingTable)
	if err != nil {
		log.Fatal(err)
	}
}
