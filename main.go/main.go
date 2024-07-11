package main

import (
	"employee-asset-mapping/db"
	"employee-asset-mapping/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB("user=youruser dbname=yourdb sslmode=disable")

	r := mux.NewRouter()

	// Employee Routes
	r.HandleFunc("/employee/createemployee", handlers.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/editemployee/{employeeId}", handlers.EditEmployee).Methods("PUT")
	r.HandleFunc("/employee/deleteemployee/{employeeId}", handlers.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/employee/{employeeId}", handlers.GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", handlers.GetAllEmployees).Methods("GET")

	// Asset Routes
	r.HandleFunc("/asset/createasset", handlers.CreateAsset).Methods("POST")
	r.HandleFunc("/asset/editasset/{assetid}", handlers.EditAsset).Methods("PUT")
	r.HandleFunc("/asset/getallasset", handlers.GetAllAssets).Methods("GET")
	r.HandleFunc("/asset/{assetId}", handlers.GetAssetById).Methods("GET")
	r.HandleFunc("/asset/deleteAsset/{assetId}", handlers.DeleteAsset).Methods("DELETE")

	// Mapping Routes
	r.HandleFunc("/mapping/assignassetmapping", handlers.AssignAssetMapping).Methods("POST")
	r.HandleFunc("/mapping/getallassets/{employeeId}", handlers.GetAllAssetsMappedToEmployee).Methods("GET")
	r.HandleFunc("/mapping/removeassetmapping/{mappingId}", handlers.RemoveAssetMapping).Methods("DELETE")

	// Dashboard Route
	r.HandleFunc("/dashboard", handlers.GetDashboard).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
