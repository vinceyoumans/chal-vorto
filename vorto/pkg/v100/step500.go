package v100

import (
	"vorto03/vorto/pkg/strucs"
)

func Step500_BuildRoute(P320 []strucs.PMap210_ROUTE) strucs.VortoMatrix500 {

	// =================================
	// step 100 build VortoMatrix
	vm500 := strucs.VortoMatrix500{}
	vm500.Max_number_of_drivers = len(P320) - 1
	vm500.MaxDriveTime_Min = 12 * 60 * 60 // 12 hours * 60 Minutes * 60 seconds

	const max_Driving_Time = 12 * 60

	for i := 1; i <= vm500.Max_number_of_drivers; i++ {

		tBMCount := strucs.DMCount{}
		tBMCount.A_Number_of_Drivers = i
		tBMCount.B_Total_NumberOfDrivenMinutes = 0.0
		tBMCount.C_Total_Cost = 0.0
		tBMCount.R_DriverMissions = []strucs.DriverMission{}
		tBMCount.Violates_12HourRule = false
		tBMCount.ZZ_debug_Total_DrivingTime = 0.0
		vm500.DMCounts = append(vm500.DMCounts, tBMCount)
	}

	// =================================
	// Step 200
	// B_Total_NumberOfDrivenMinutes float32 `B_Total_Number_Of_Driven_Minutes`
	totalMinutes := step200_Compute_TotalNumberOfDrivenHours(P320, max_Driving_Time)

	for iDM, _ := range vm500.DMCounts {
		vm500.DMCounts[iDM].B_Total_NumberOfDrivenMinutes = totalMinutes
		vm500.DMCounts[iDM].B_Average_DriveTime = totalMinutes / float64(vm500.DMCounts[iDM].A_Number_of_Drivers)
	}
	vm500.B_Total_NumberOfDrivenMinutes = totalMinutes

	return vm500

}

// ==================================================
// func step200_Compute_TotalNumberOfDrivenHours(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time int) float64 {
func step200_Compute_TotalNumberOfDrivenHours(P320 []strucs.PMap210_ROUTE, max_Driving_Time int) float64 {
	total_NumberOfDrivenMinutes := 0.0
	for iPM320, valP320 := range P320 {
		if iPM320 == 0 {
			continue
		}

		if iPM320 <= len(P320) {
			total_NumberOfDrivenMinutes += valP320.ABC_RouteDistance
		} else {
			total_NumberOfDrivenMinutes += valP320.ABCDepot_RouteDistance
		}
	}
	return total_NumberOfDrivenMinutes
}

// ==================================================
// step300_checkForViolations
func step300B_CalcDriversMission(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time float64) strucs.DMCount {

	// numDrivers := valDM.A_Number_of_Drivers
	// currentDriver := 1

	return valDM
}

// ==================================================
// step300_checkForViolations
func step300_checkForViolations(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time float64) strucs.DMCount {

	tempDMCount := valDM
	numOfDrivers := tempDMCount.A_Number_of_Drivers
	currentDriver := 1
	currentDriver_RunningDriveTime := 0.0

	tLoadNumber := []int{}
	mapDriverMission := make(map[int]strucs.DriverMission, numOfDrivers)

	for iP320, valP320 := range P320 {
		if iP320 == 0 {
			continue
		}

		test100 := (currentDriver_RunningDriveTime + valP320.ABC_RouteDistance) < max_Driving_Time
		test110 := (currentDriver_RunningDriveTime + valP320.ABCDepot_RouteDistance) < max_Driving_Time
		test200 := currentDriver <= numOfDrivers

		if !test200 {
			// To Many Drivers
			tempDMCount.Violates_12HourRule = true
		}

		if !test110 {
			// Adding this route to DriverRoute will exceed 12 hour rule
			tempDMCount.Violates_12HourRule = true
			tLoadNumber = append(tLoadNumber, valP320.Current_LoadNumber)
			mapDriverMission[currentDriver] = strucs.DriverMission{
				LoadNumberS:      tLoadNumber,
				RunningDriveTime: currentDriver_RunningDriveTime,
			}
			currentDriver++
			currentDriver_RunningDriveTime = 0.0
			tLoadNumber = []int{}
		}

		if test100 && test200 {
			// CurrentDriver not Violating 12hour rule
			currentDriver_RunningDriveTime += valP320.ABC_RouteDistance
			tLoadNumber = append(tLoadNumber, valP320.Current_LoadNumber)
		}

	}

	// // ==================================================
	// // step300_checkForViolations
	// // func step300_checkForViolations(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time float64) strucs.DMCount {
	// func step300_checkForViolations(valDM strucs.DMCount, max_Driving_Time float64) strucs.DMCount {
	// 	tempDMCount := valDM
	// 	numOfDrivers := tempDMCount.A_Number_of_Drivers
	// 	currentDriver := 1
	// 	currentDriver_RunningDriveTime := 0.0

	// 	tLoadNumber := []int{}
	// 	mapDriverMission := make(map[int]strucs.DriverMission, numOfDrivers)

	// 	for iP320, valP320 := range P320 {
	// 		if iP320 == 0 {
	// 			continue
	// 		}

	// 		test100 := (currentDriver_RunningDriveTime + valP320.ABC_RouteDistance) < max_Driving_Time
	// 		test110 := (currentDriver_RunningDriveTime + valP320.ABCDepot_RouteDistance) < max_Driving_Time
	// 		test200 := currentDriver <= numOfDrivers

	// 		if test100 && test110 && test200 {
	// 			// CurrentDriver not Violating 12hour rule
	// 			currentDriver_RunningDriveTime += valP320.ABC_RouteDistance
	// 			tLoadNumber = append(tLoadNumber, valP320.Current_LoadNumber)
	// 			continue
	// 		}

	// 		if !test110 {
	// 			// CurrentDriver has hit the max driving time
	// 			// Including returing to Depot
	// 			tempDMCount.Violates_12HourRule = true
	// 			tLoadNumber = append(tLoadNumber, valP320.Current_LoadNumber)
	// 			mapDriverMission[currentDriver] = strucs.DriverMission{
	// 				LoadNumber:       tLoadNumber,
	// 				RunningDriveTime: currentDriver_RunningDriveTime,
	// 			}
	// 			currentDriver++
	// 			currentDriver_RunningDriveTime = 0.0
	// 			tLoadNumber = []int{}
	// 		}

	// 	}

	return tempDMCount
}

func step300_BuildDMCounts(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time int) strucs.DMCount {

	tempDMCount := valDM
	// driverMap := make(map[int]strucs.DriverMission, tempDMCount.A_Number_of_Drivers)
	// currentDriver := 1
	// CurrentDriver_totalDriveTime := 0.0
	// CurrentTimePDepotreturn := 0.0 //CurrentRunTIme Plus return to Depot

	// for iDM2, valP320 := range P320 {
	// 	if iDM2 == 0 {
	// 		continue
	// 	}

	// driverMap[iDM2] = strucs.DriverMission{
	// 	LoadNumber       []int   `json:"max_Number_of_Drivers"`
	// 	RunningDriveTime float64 `json:"DMCounts"`
	// }

	// CurrentTimePDepotreturn = CurrentDriver_totalDriveTime + valP320.ABCDepot_RouteDistance
	// test01 := (CurrentTimePDepotreturn > max_Driving_Time)
	// // test02 := ()

	// if CurrentTimePDepotreturn > max_Driving_Time {
	// 	// CurrentDriver has hit the max driving time
	// 	tLoadNumber := valP320.Current_LoadNumber
	// 	driverMap[currentDriver].LoadNumber = append(driverMap[currentDriver].LoadNumber, tLoadNumber)

	// }

	// 	currentDriver++
	// 	CurrentDriver_totalDriveTime = 0.0
	// }

	return tempDMCount
}

// =================================
// perhaps over kill and probably not correct
// But want to confirm there are enought Drivers
// func step200_checkForViolations(valDM strucs.DMCount, P320 []strucs.PMap210_ROUTE, max_Driving_Time int) strucs.DMCount {

// 	tDC := valDM
// 	debug_totalDriveTime := 0.0
// 	driverRunningTime := 0.0
// 	driver := 1
// 	for iDM2, valP320 := range P320 {
// 		if iDM2 == 0 {
// 			continue
// 		}

// 		if (driverRunningTime + valP320.ABC_RouteDistance) > float64(max_Driving_Time) {
// 			if driver < tDC.A_Number_of_Drivers {
// 				tDC.Violates_12HourRule = true
// 				break
// 			}
// 			tDC.Violates_12HourRule = true
// 			break
// 		}

// 	}
// 	return tDC
// }
