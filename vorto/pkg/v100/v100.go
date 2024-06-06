package v100

import (
	slogpkg "github.com/vinceyoumans/chal-vorto/vorto/pkg/slogPkg"
	"github.com/vinceyoumans/chal-vorto/vorto/pkg/strucs"
	util "github.com/vinceyoumans/chal-vorto/vorto/pkg/util"
)

func V100Start(pp string) [][]int {

	// Prob100 := DigestProblemFile("../training", pp)
	Prob100 := DigestProblemFile("", pp)

	util.PP100SaveProblemJSON("../output/Problems/prob20", "", "prob20.json", Prob100)
	slogpkg.LogVortoP100("Test LogMessage")

	P200 := BuildP1000(Prob100)
	// util.PP200SaveProblemJSON("../output/P200", "", "p200.json", P200)

	// Just creating
	// P300, Prob300 := Step300_CreateRouteSortA(P200, Prob100)
	P300, _ := Step300_CreateRouteSortA(P200, Prob100)

	// util.PP200SaveProblemJSON("../output/P300", "", "p300.json", P300)
	// util.PP100SaveProblemJSON("../output/Problems/prob300", "", "prob300.json", Prob300)

	// Update P300 to reflect Fields for Computation.
	P320 := Step320_UpdateFields(P300, Prob100)
	util.PP200SaveProblemJSON("../output/problems/P320", "", "p320.json", P320)

	// Step 500 - Build Route Matrix
	PM500 := Step500_BuildRoute(P320)
	util.PP500SaveProblemJSON("../output/problems/PM500", "", "pm500.json", PM500)

	// Step 520 - BuildDriverMissions
	// This version is flawed.
	// PM520 := Step520A_BuildRoute(PM500, P320)
	// util.PP500SaveProblemJSON("../output/problems/PM520", "", "pm520.json", PM520)

	PM530 := Step530_BuildRoute(PM500, P320)
	util.PP500SaveProblemJSON("../output/problems/PM530", "", "pm530.json", PM530)

	res1000 := Step1000(PM530)
	return res1000
}

func Step1000(PM530 strucs.VortoMatrix500) [][]int {
	r1000 := [][]int{}
	for _, valDriver := range PM530.DMCounts[0].R_DriverMissions {
		// fmt.Println("DriverID", iDriver, "Mission", valDriver)
		r1000 = append(r1000, valDriver.LoadNumberS)
	}
	return r1000
}

func Step530_BuildRoute(PM530 strucs.VortoMatrix500, P320 []strucs.PMap210_ROUTE) strucs.VortoMatrix500 {
	// fmt.Println("======.  Step530 C=======")
	// fmt.Println(PM530)
	const maxDrive_Time_Minutes = 12 * 60

	// MaxRoutes := len(P320)

	// r_driversMissions := strucs.DriverMission{}

	tDriverID := 1
	tRunningDriveTime := 0.0
	tDriverMission := strucs.DriverMission{}
	tsiDriverMission := []int{}
	mDriverMissionS := make(map[int]strucs.DriverMission)

	for iP320, valP320 := range P320 {
		if iP320 == 0 {
			//Skip LoadZero
			continue
		}

		// fmt.Println("======.  Step530 D=======", iP320)
		// fmt.Println("======.  Step530 E=======", valP320)

		temp_RunningDriveTime_ABC := tRunningDriveTime + P320[iP320].ABC_RouteDistance
		temp_RunningDriveTime_ABCDepo := tRunningDriveTime + P320[iP320].ABCDepot_RouteDistance

		// If Driver Accepts load, will Violate 12Hour Rule
		test100 := (temp_RunningDriveTime_ABC > maxDrive_Time_Minutes)
		test200 := (temp_RunningDriveTime_ABCDepo > maxDrive_Time_Minutes)

		if test100 || test200 {
			// This Driver should return to depot

			// fmt.Println("======.  Step530 F=======", iP320)
			// fmt.Println("======.  Step530 G=======", valP320)

			tDriverMission.DriverID = tDriverID
			tDriverMission.RunningDriveTime = temp_RunningDriveTime_ABCDepo
			tDriverMission.LoadNumberS = tsiDriverMission // tempSlice of Integers

			mDriverMissionS[tDriverID] = tDriverMission

			// reset the counters
			tDriverID++
			tRunningDriveTime = 0.0
			tsiDriverMission = []int{}
			continue
		}

		tRunningDriveTime = temp_RunningDriveTime_ABC
		tsiDriverMission = append(tsiDriverMission, valP320.Current_LoadNumber)
	}

	tDMCount := strucs.DMCount{}
	tDMCount.A_Number_of_Drivers = len(mDriverMissionS)

	// this is redundant
	tDMCount.B_Total_NumberOfDrivenMinutes = PM530.B_Total_NumberOfDrivenMinutes

	//total_cost = 500*number_of_drivers + total_number_of_driven_minutes
	tDMCount.C_Total_Cost = Step500_GetTotalCost(tDMCount.A_Number_of_Drivers, tDMCount.B_Total_NumberOfDrivenMinutes)

	// for i, v := range mDriverMissionS {
	// 	tDMCount.R_DriverMissions = append(tDMCount.R_DriverMissions, v)
	// 	fmt.Println(i, v)
	// }

	for i := 1; i <= len(mDriverMissionS); i++ {
		tDMCount.R_DriverMissions = append(tDMCount.R_DriverMissions, mDriverMissionS[i])
		// fmt.Println(i, mDriverMissionS[i])
	}
	// fmt.Println("======.  Step530 H=======")
	// fmt.Println(tDMCount)
	PM530.DMCounts = []strucs.DMCount{}
	PM530.DMCounts = append(PM530.DMCounts, tDMCount)
	return PM530

}

func Step500_GetTotalCost(number_of_drivers int, total_number_of_driven_minutes float64) float64 {
	// Will return the TotalCost, which is what the assignment was called to do
	total_cost := float64(500*number_of_drivers) + total_number_of_driven_minutes
	return total_cost
}

// ==================================================
// Step520_BuildRoute - will populate the driversMission
func Step520A_BuildRoute(PM500 strucs.VortoMatrix500, P320 []strucs.PMap210_ROUTE) strucs.VortoMatrix500 {

	for iDM, valDM := range PM500.DMCounts {

		// fmt.Println("======.  Step520 B=======", iDM)
		// fmt.Println("======.  Step520 C=======", valDM)

		NumberOfDrivers := valDM.A_Number_of_Drivers

		tvalDM := S520A_100(valDM, P320, NumberOfDrivers, iDM)
		// fmt.Println("tvalDM", tvalDM)

		PM500.DMCounts[iDM] = tvalDM

	}
	return PM500

}

func S520A_100(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, NumberOfDrivers int, iDM int) strucs.DMCount {

	// fmt.Println("======.  Step520 S520A_100 AA=======", valDM)

	const maxDriveTime_Seconds = 12 * 60 * 60 // 12 hours * 60 Minutes * 60 seconds
	const maxDrive_Time_Minutes = 12 * 60

	mapValDM := make(map[int]strucs.DriverMission, NumberOfDrivers)

	maxRN := len(P320) - 1 // Max RouteNumbers
	ttLNS := []int{}
	//tempDC := strucs.DriverCount{}

	ttRunningDriveTime := 0.0
	temp_RunningDriveTime_ABC := 0.0
	temp_RunningDriveTime_ABCDepo := 0.0
	lastRoute := 1

	ZZ_ToMany_Drivers := false
	ZZ_ToFewDrivers := false
	ZZ_Perfect_Drivers := false

	for iDriver := 1; iDriver <= NumberOfDrivers; iDriver++ {

		// fmt.Println("======.  Step520 S520A_100. BB======iDriver=", iDriver)
		// fmt.Println("======.  Step520 S520A_100. BB======NumberOfDrivers=", NumberOfDrivers)

		for iLR := lastRoute; iLR <= maxRN; iLR++ {

			// fmt.Println("======.  Step520 S520A_100. CC======iLR=", iLR)
			// fmt.Println("======.  Step520 S520A_100. CC======lastRoute=", lastRoute)
			// fmt.Println("======.  Step520 S520A_100. CC======maxRN=", maxRN)

			temp_RunningDriveTime_ABC = ttRunningDriveTime + P320[iLR].ABC_RouteDistance
			temp_RunningDriveTime_ABCDepo = ttRunningDriveTime + P320[iLR].ABCDepot_RouteDistance
			// test100 - confirms driver is not violating 12 hour rule
			//		LOGIC: Should Never Happen
			// test100 := (RunningDriveTime >= maxDrive_Time_Minutes)

			// test200 - Driver WILL Exceed 12 hours if Run This Load
			test200 := (temp_RunningDriveTime_ABC >= maxDrive_Time_Minutes)

			// test300 - Driver WILL exceed 12rule if return to depot
			//		LOGIC: Depot return could be shorter or Longer than the next route
			//		A return from this point will ALWAYS be shorter than Next Route + depotReturn
			test300 := (temp_RunningDriveTime_ABCDepo >= maxDriveTime_Seconds)

			// test600 - This is the optimal number of drivers
			test600 := (iLR == maxRN)

			if test200 || test300 {
				// This driver needs to return Home BEFORE accepting this Load
				lastRoute = iLR
				break
			}

			if test600 {
				ZZ_Perfect_Drivers = true
			}

			// Else - Driver may continue accepting Loads with out exceeding 12 hour rule
			ttLNS = append(ttLNS, P320[iLR].Current_LoadNumber)
			ttRunningDriveTime = temp_RunningDriveTime_ABC

			temp_mapValDM := strucs.DriverMission{
				DriverID:         iDriver,
				LoadNumberS:      ttLNS,              // []int
				RunningDriveTime: ttRunningDriveTime, //float64
			}
			mapValDM[iDriver] = temp_mapValDM

		}

		// Record this Driver
		tDM := strucs.DriverMission{

			DriverID:         mapValDM[iDriver].DriverID,         // []int
			LoadNumberS:      mapValDM[iDriver].LoadNumberS,      // []int
			RunningDriveTime: mapValDM[iDriver].RunningDriveTime, //float64
		}
		valDM.R_DriverMissions = append(valDM.R_DriverMissions, tDM)
		ttLNS = []int{}
		ttRunningDriveTime = 0.0

		if ZZ_Perfect_Drivers {
			// No point in continuing with buildOut
			break
		}
	}

	valDM.ZZ_Perfect_Drivers = ZZ_Perfect_Drivers
	valDM.ZZ_ToFewDrivers = ZZ_ToFewDrivers
	valDM.ZZ_ToMany_Drivers = ZZ_ToMany_Drivers

	return valDM
}

// func S520B_100(valDM strucs.DMCount, Prob100 []strucs.Problem100, i, iDM int) {
// 	panic("unimplemented")
// }
