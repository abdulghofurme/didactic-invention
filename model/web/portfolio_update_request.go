package web

type PortfolioUpdateRequest struct {
	ID          string
	Name        string
	Description string
	Balance     int
	Nominal     int
}
