package web

type PortfolioResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Balance     int    `json:"balance"`
	Nominal     int    `json:"nominal"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"delete_at"`
}
