package util

import (
	"path/filepath"

	strucs "github.com/vinceyoumans/chal-vorto/vorto/pkg/strucs"
)

func CreateJSONDir(baseDir, SubDir string) string {
	MakeDir(baseDir)
	SD := filepath.Join(baseDir, SubDir)
	MakeDir(SD)
	return SD
}

// ====================================
func PP100SaveProblemJSON(pathToProblemJSON, SaveDir, JsoneFileName string, JTSS []strucs.Problem100) {
	//--Step 30 -------------------------
	// Save Struct for review
	pathToPJSON := CreateJSONDir(pathToProblemJSON, SaveDir)

	PP100(pathToPJSON, JsoneFileName, JTSS)
}

// ====================================
func PP200SaveProblemJSON(pathToProblemJSON, SaveDir, JsoneFileName string, JTSS []strucs.PMap210_ROUTE) {
	//--Step 30 -------------------------
	// Save Struct for review
	pathToPJSON := CreateJSONDir(pathToProblemJSON, SaveDir)

	PP200(pathToPJSON, JsoneFileName, JTSS)
}

// =======================================
func PP500SaveProblemJSON(pathToProblemJSON, SaveDir, JsoneFileName string, JTSS strucs.VortoMatrix500) {
	pathToPJSON := CreateJSONDir(pathToProblemJSON, SaveDir)
	PP500(pathToPJSON, JsoneFileName, JTSS)
}
