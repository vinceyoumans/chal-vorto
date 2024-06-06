package util

import (
	"log"
	"os"
)

// MakeDir - Creates Local directories for saving logs and output
func MakeDir(baseDir string) {
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		err := os.MkdirAll(baseDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %s", err)
		}
	}
}

// MakeOutputDirs - Creates Local directories for saving logs and output
func MakeOutputDirs() {
	outputDir := "../output"
	MakeDir(outputDir)

	outputDir = "../output/ret"
	MakeDir(outputDir)

	outputDir = "../output/slog"
	MakeDir(outputDir)

	outputDir = "../output/strucs"
	MakeDir(outputDir)

}
