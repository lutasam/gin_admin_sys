package vo

type CommodityVO struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Count int     `json:"count"`
}
