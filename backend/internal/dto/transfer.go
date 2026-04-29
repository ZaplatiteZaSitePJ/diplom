package dto

type MoveItemRequest struct {
	ItemID string  `json:"item_id"`
	ToType string  `json:"to_type"`
	ToID   *string `json:"to_id"`
}