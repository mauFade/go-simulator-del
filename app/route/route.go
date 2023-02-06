package route

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Latitude  float64
	Longitude float64
}
