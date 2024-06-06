package strucs

// --------------------------------
type Problem100 struct {
	LoadNumber int
	Pickup     LatLong
	DropOff    LatLong
	PickedUp   bool
}
type LatLong struct {
	Latitude  float64
	Longitude float64
}
