package strucs

// depot located at (0,0)
var depot = LatLong{
	Latitude:  0.0,
	Longitude: 0.0,
}

// Depot returns a copy of the depot's coordinates to ensure immutability
func Depot() LatLong {
	return depot
}

const MaxFloat = float64(1<<63 - 1) // Largest possible float64 value
