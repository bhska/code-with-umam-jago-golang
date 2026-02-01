package entity

type Product struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	Harga      int     `json:"harga"`
	CategoryID int     `json:"category_id"`
	Category   *Category `json:"category,omitempty"`
}
