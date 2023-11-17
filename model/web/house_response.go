package web

type HouseResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BlockName   string `json:"block_name"`
	BlockNumber int    `json:"block_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"delete_at"`
}
