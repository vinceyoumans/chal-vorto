package util

import (
	"encoding/json"
	"fmt"
	"os"
	strucs "vorto03/vorto/pkg/strucs"
)

// PrettyPrint

func PP100(SavePath, JSONFileName string, probStruc []strucs.Problem100) {

	jsonOutput, err := json.MarshalIndent(probStruc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	outputDir := SavePath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	fileJSON, err := os.Create(SavePath + "/" + JSONFileName)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer fileJSON.Close()

	_, err = fileJSON.Write(jsonOutput)
	if err != nil {
		fmt.Println(err)
		// return nill
	}

}

func PP200(SavePath, JSONFileName string, probStruc []strucs.PMap210_ROUTE) {

	jsonOutput, err := json.MarshalIndent(probStruc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	outputDir := SavePath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	fileJSON, err := os.Create(SavePath + "/" + JSONFileName)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer fileJSON.Close()

	_, err = fileJSON.Write(jsonOutput)
	if err != nil {
		fmt.Println(err)
		// return nill
	}

}

func PP500(SavePath, JSONFileName string, probStruc strucs.VortoMatrix500) {

	jsonOutput, err := json.MarshalIndent(probStruc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	outputDir := SavePath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	fileJSON, err := os.Create(SavePath + "/" + JSONFileName)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer fileJSON.Close()

	_, err = fileJSON.Write(jsonOutput)
	if err != nil {
		fmt.Println(err)
		// return nill
	}

}
