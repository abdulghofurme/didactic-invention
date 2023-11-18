package web

type PortfolioCreateRequest struct {
	Name        string
	Description string
	Balance     int
	Nominal     int
}
