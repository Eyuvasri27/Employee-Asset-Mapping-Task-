package models

import (
	"github.com/google/uuid"
)

type Asset struct {
	AssetId   uuid.UUID `json:"asset_id"`
	AssetName string    `json:"asset_name"`
	AssetType string    `json:"asset_type"`
}
