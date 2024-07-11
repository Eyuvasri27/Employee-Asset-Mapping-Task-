package models

import (
	"github.com/google/uuid"
)

type EmployeeAssetMapping struct {
    ID     uuid.UUID `json:"id"`
    EmpId  uuid.UUID `json:"emp_id"`
    AssetId uuid.UUID `json:"asset_id"`
}
