package handlers

import (
	"database/sql"
	"employee-asset-mapping/db"
	"employee-asset-mapping/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	asset.AssetId = uuid.New()
	_, err := db.DB.Exec(`INSERT INTO assets (asset_id, asset_name, asset_type) VALUES ($1, $2, $3)`, asset.AssetId, asset.AssetName, asset.AssetType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asset)
}

func EditAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetId, err := uuid.Parse(vars["assetid"])
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = db.DB.Exec(`UPDATE assets SET asset_name=$1, asset_type=$2 WHERE asset_id=$3`, asset.AssetName, asset.AssetType, assetId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asset)
}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetId, err := uuid.Parse(vars["assetId"])
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	_, err = db.DB.Exec(`DELETE FROM assets WHERE asset_id=$1`, assetId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetAssetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetId, err := uuid.Parse(vars["assetId"])
	if err != nil {
		http.Error(w, "Invalid asset ID", http.StatusBadRequest)
		return
	}
	var asset models.Asset
	err = db.DB.QueryRow(`SELECT asset_id, asset_name, asset_type FROM assets WHERE asset_id=$1`, assetId).Scan(&asset.AssetId, &asset.AssetName, &asset.AssetType)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Asset not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(asset)
}

func GetAllAssets(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT asset_id, asset_name, asset_type FROM assets`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var assets []models.Asset
	for rows.Next() {
		var asset models.Asset
		err := rows.Scan(&asset.AssetId, &asset.AssetName, &asset.AssetType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		assets = append(assets, asset)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(assets)
}
