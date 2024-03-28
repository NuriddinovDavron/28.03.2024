package models

// Product ...
type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
	Price       int    `json:"price"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

type ProductId struct {
	Id string `json:"id"`
}
