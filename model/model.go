package model

type InventoryInfo struct {
	UniqueItemId      string  `json:"unique_item_id"`
	ItemName          string  `json:"item_name"`
	SupplierId        string  `json:"supplier_id"`
	IsAvailable       bool    `json:"is_available"`
	AvailableQuantity int     `json:"available_quantity"`
	PricePerUnit      float32 `json:"price_per_unit"`
}
