package bo

type HelloRequest struct {
	Username string `json:"username"`
}

/*
=================================================
  - Login Module

==================================================
*/
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"-"`
	Avatar   string `json:"avatar" binding:"-"`
	Sign     string `json:"sign" binding:"-"`
}

type ActiveRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
}

/*
=================================================
  - Commodity Module

==================================================
*/
type AddCommodityRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Count int     `json:"count" binding:"required"`
}

type UpdateCommodityRequest struct {
	ID    string  `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Count int     `json:"count" binding:"required"`
}

type FindAllCommoditiesRequest struct{}

type DeleteCommodityRequest struct {
	ID string `json:"id" binding:"required"`
}
