package v100

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	slogpkg "github.com/vinceyoumans/chal-vorto/vorto/pkg/slogPkg"

	strucs "github.com/vinceyoumans/chal-vorto/vorto/pkg/strucs"
)

// DigestProblemFile - opens the ProblemFile and returns Struc
// BaseDir - the directory where the problem file is located
// FileToOpen - Problem file to be studied
func DigestProblemFile(BaseDir, FileToOpen string) []strucs.Problem100 {
	// func DigestProblemFile(BaseDir, FileToOpen, SaveDir, JSONToSave string) []strucs.Problem100 {
	// fmt.Println("----   in DigestProblemFile")

	//--Step 10 -------------------------
	// Seed Everything with an array 0
	var probStrucS []strucs.Problem100
	depotLocation := strucs.Problem100{
		LoadNumber: 0,
		Pickup: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
		DropOff: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
	}
	probStrucS = append(probStrucS, depotLocation)

	//--Step 20 -------------------------
	// Open and Digest Problem File
	pathToProblem := filepath.Join(BaseDir, FileToOpen)
	file, err := os.Open(pathToProblem)
	if err != nil {
		fmt.Println("Error opening file:", err)
		// return nil
	}
	defer file.Close()

	// var probStrucS []struc.Problem100
	scanner := bufio.NewScanner(file)
	// Skip the header line
	if scanner.Scan() {
		header := scanner.Text()
		// fmt.Println("Header:", header)
		slogpkg.LogVortoP100(header)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		// str = strings.Trim(str, "()")
		parts[1] = strings.Trim(parts[1], "()")
		parts[2] = strings.Trim(parts[2], "()")

		if len(parts) != 3 {
			fmt.Println("*****  Invalid line format:", line)
			continue
		}

		loadNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing load number:", err)
			continue
		}
		pickup := parseLatLong(parts[1])
		dropoff := parseLatLong(parts[2])

		probStruc := strucs.Problem100{
			LoadNumber: loadNumber,
			Pickup:     pickup,
			DropOff:    dropoff,
		}

		probStrucS = append(probStrucS, probStruc)

	}

	return probStrucS
}

func parseLatLong(str string) strucs.LatLong {
	str = strings.Trim(str, "()")
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		fmt.Println("Invalid LatLong format:", str)
		return strucs.LatLong{}
	}

	latitude, err1 := strconv.ParseFloat(parts[0], 64)
	longitude, err2 := strconv.ParseFloat(parts[1], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing LatLong:", err1, err2)
		return strucs.LatLong{}
	}

	return strucs.LatLong{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
