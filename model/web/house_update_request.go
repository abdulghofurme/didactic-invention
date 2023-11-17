package web

type HouseUpdateRequest struct {
	ID          string
	BlockName   string `json:"block_name"`
	BlockNumber int    `json:"block_number,string"`
}
