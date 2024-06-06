package strucs

type VortoMatrix500 struct {
	Max_number_of_drivers         int
	MaxDriveTime_Min              float64
	B_Total_NumberOfDrivenMinutes float64
	DMCounts                      []DMCount
}

// total_cost = 500*number_of_drivers + total_number_of_driven_minutes

type DMCount struct {
	A_Number_of_Drivers           int // number of drivers
	B_Total_NumberOfDrivenMinutes float64
	B_Average_DriveTime           float64
	C_Total_Cost                  float64
	R_DriverMissions              []DriverMission
	Violates_12HourRule           bool
	ZZ_debug_Total_DrivingTime    float64
	ZZ_ToMany_Drivers             bool
	ZZ_ToFewDrivers               bool
	ZZ_Perfect_Drivers            bool
}

type DriverMission struct {
	DriverID         int
	LoadNumberS      []int
	RunningDriveTime float64
}

// total_cost = 500*number_of_drivers + total_number_of_driven_minutes
