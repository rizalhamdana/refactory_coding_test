package model

type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

type Inventory struct {
	InventoryID int       `json:"inventory_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	PurchasedAt int       `json:"purchased_at"`
	Placement   Placement `json:"placement"`
}
