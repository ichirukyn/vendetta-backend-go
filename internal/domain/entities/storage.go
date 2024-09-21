package entities

type Storage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Count       int    `json:"count"`
	IsInventory bool   `json:"is_inventory"`
	CellLimit   int    `json:"cell_limit"`
	LockItemID  string `json:"lock_item_id"`
	IsLocked    bool   `json:"is_locked"`
}

type StorageItem struct {
	ID         string `json:"id"`
	StorageID  string `json:"storage_id"`
	ItemID     string `json:"item_id"`
	Count      int    `json:"count"`
	IsStack    bool   `json:"is_stack"`
	IsTransfer bool   `json:"is_transfer"`

	Storage Storage `json:"storage"`
}
