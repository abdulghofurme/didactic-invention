package web

type HouseCreateRequest struct {
	BlockName   string `json:"block_name"`
	BlockNumber int    `json:"block_number,string"`
}
