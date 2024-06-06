package v100

import (
	slogpkg "github.com/vinceyoumans/chal-vorto/vorto/pkg/slogPkg"
	strucs "github.com/vinceyoumans/chal-vorto/vorto/pkg/strucs"
)

func BuildP1000(PM100 []strucs.Problem100) []strucs.PMap210_ROUTE {
	slogpkg.LogVortoP100("starting BuildP1000")

	var PM1000 strucs.PMap210_ROUTE
	var PM1000S []strucs.PMap210_ROUTE

	for i := 0; i < len(PM100); i++ {
		PM1000.RouteID = i
		PM1000.PREV_LoadNumber = 0
		PM1000.Current_LoadNumber = 0

		PM1000.A_StartLoaction = strucs.Depot()
		PM1000.B_PickUpLatLong = strucs.Depot()
		PM1000.C_DropOffLatLong = strucs.Depot()
		PM1000.AB_RouteDistance = strucs.MaxFloat
		PM1000.BC_RouteDistance = strucs.MaxFloat
		PM1000.ABC_RouteDistance = strucs.MaxFloat
		PM1000S = append(PM1000S, PM1000)
	}
	return PM1000S
}
