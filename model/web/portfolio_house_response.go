package web

type PortfolioHouseResponse struct {
	ID        string            `json:"id"`
	Portfolio PortfolioResponse `json:"portfolio"`
	House     HouseResponse     `json:"house"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	DeletedAt string            `json:"delete_at"`
}
