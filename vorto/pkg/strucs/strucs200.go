package strucs

type PMap210_ROUTE struct {
	RouteID                int `json:"RouteID"`
	PREV_LoadNumber        int `json:"prev_loadNumber"`
	Current_LoadNumber     int `json:"current_loadNumber"`
	A_StartLoaction        LatLong
	B_PickUpLatLong        LatLong
	C_DropOffLatLong       LatLong
	AB_RouteDistance       float64 `json:"AB_RouteDistance"`
	BC_RouteDistance       float64 `json:"BC_RouteDistance"`
	ABC_RouteDistance      float64 `json:"ABC_RouteDistance"`
	CDepot_RouteDistance   float64 `json:"CDepot_RouteDistance"`
	ABCDepot_RouteDistance float64 `json:"ABCDepot_RouteDistance"`
}
