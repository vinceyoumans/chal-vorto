package v100

import (
	"math"

	"github.com/vinceyoumans/chal-vorto/vorto/pkg/strucs"
)

// Step300_CreateRouteSortA - builds a sorted list of routes assuming Clostest PickUp is assigend Next route
func Step300_CreateRouteSortA(p300 []strucs.PMap210_ROUTE, p100 []strucs.Problem100) ([]strucs.PMap210_ROUTE, []strucs.Problem100) {

	// for iP300, valP300 := range p300 {
	for iP300, _ := range p300 {
		// fmt.Println("iP300 - valP300", iP300, valP300)
		if iP300 == 0 {
			// Skip the first RouteID
			continue
		}

		// slogpkg.LogVortoP100("ip300 - valP300")
		// jsonData, err := json.Marshal(iP300)
		// if err != nil {
		// 	fmt.Println("Error marshalling to JSON:", err)
		// }
		// jsonString := string(jsonData)
		// slogpkg.LogVortoP100(jsonString)

		// Assumes starting at 0,0 - Depot.
		// Last Known Position is from previous route
		// Assumes starting at 0,0 - Depot.
		// Last Known Position is from previous route.
		previousLoadNumber := p300[iP300-1].Current_LoadNumber
		assignedLoadNumber := getAssignedLoadNumber(previousLoadNumber, &p100)

		// mark p100 as Pickedup
		p100[assignedLoadNumber].PickedUp = true
		p300[iP300].Current_LoadNumber = assignedLoadNumber

	}
	return p300, p100

}

func Step320_UpdateFields(p320 []strucs.PMap210_ROUTE, p100 []strucs.Problem100) []strucs.PMap210_ROUTE {

	for iP, valp320 := range p320 {

		if iP == 0 {
			// Skip the first RouteID
			continue
		}

		current_LoadNumber_Assignment := valp320.Current_LoadNumber

		valp320.PREV_LoadNumber = p320[iP-1].Current_LoadNumber
		valp320.A_StartLoaction = p320[iP-1].C_DropOffLatLong

		valp320.B_PickUpLatLong = p100[current_LoadNumber_Assignment].Pickup
		valp320.C_DropOffLatLong = p100[current_LoadNumber_Assignment].DropOff

		p320[iP] = valp320

	}

	for iP, valp320 := range p320 {

		if iP == 0 {
			// Skip the first RouteID
			continue
		}

		valp320.AB_RouteDistance = getEuclideanDistance300(valp320.A_StartLoaction, valp320.B_PickUpLatLong)
		valp320.BC_RouteDistance = getEuclideanDistance300(valp320.B_PickUpLatLong, valp320.C_DropOffLatLong)
		valp320.CDepot_RouteDistance = getEuclideanDistance300(valp320.C_DropOffLatLong, strucs.Depot())
		valp320.ABC_RouteDistance = valp320.AB_RouteDistance + valp320.BC_RouteDistance
		valp320.ABCDepot_RouteDistance = valp320.ABC_RouteDistance + valp320.CDepot_RouteDistance

		// fmt.Println("iP - valp320", iP, valp320)
		p320[iP] = valp320

	}

	return p320
}

// ==========================================================
// func getAssignedLoadNumber(pln int, valP300 strucs.PMap210_ROUTE, problem *[]strucs.Problem100) int {
func getAssignedLoadNumber(pln int, problem *[]strucs.Problem100) int {

	// step10 - get current location from previous valP300
	lastPostition := (*problem)[pln].DropOff
	shortestDistance := float64(1<<63 - 1)
	assignedLoadNumber := 0
	totalNumberOfLoads := len(*problem) - 1

	for iP := 1; iP <= totalNumberOfLoads; iP++ {
		valP := (*problem)[iP]
		if !valP.PickedUp {
			// step11 - get distance from last known position to current location
			distance := getEuclideanDistance300(lastPostition, valP.Pickup)

			// step12 - get shortest distance
			if distance < shortestDistance {
				shortestDistance = distance
				assignedLoadNumber = iP
			}
		}

	}

	return assignedLoadNumber

}

func getEuclideanDistance300(P0_lastPostition, P1_CandidateNextPosition strucs.LatLong) float64 {
	prt_lat := P1_CandidateNextPosition.Latitude - P0_lastPostition.Latitude
	prt_lon := P1_CandidateNextPosition.Longitude - P0_lastPostition.Longitude

	return math.Sqrt(math.Pow(prt_lat, 2) + math.Pow(prt_lon, 2))

}

// Each load has a pickup location and a dropoff location,
// each specified by a Cartesian point.
// A driver completes a load by driving to the pickup location,
// picking up the load, driving to the dropoff,
// and dropping off the load.
// The time required to drive from one point to another,
// in minutes, is the Euclidean distance between them.
// That is, to drive from (x1, y1) to (x2, y2) takes sqrt((x2-x1)^2 + (y2-y1)^2) minutes.
