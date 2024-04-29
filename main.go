package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define paths
	contractFileName := "Test"
	contractFile := "./sol/" + contractFileName + ".sol" // Path to the Solidity contract
	outputDir := "./solc-output"                         // Directory to store ABI and BIN files
	apiDir := "./api"                                    // Directory to store generated Go bindings
	generatedFileName := "Meow"
	generatedPackageName := "api"

	// Ensure output and API directories exist
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}
	err = os.MkdirAll(apiDir, os.ModePerm)
	if err != nil {
		log.Fatalf(" %v\n: %s not writable", err, outputDir)
	}

	// Compile the contract using solc
	cmd := exec.Command("solc", "--optimize", "--abi", "--bin", "--overwrite", "--output-dir", outputDir, contractFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to compile contract: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Contract compiled successfully")

	// Find all ABI files
	abiFiles, err := filepath.Glob(filepath.Join(outputDir, "*.abi"))
	if err != nil {
		panic(err) // Handle error appropriately in real usage
	}

	// Find all BIN files
	binFiles, err := filepath.Glob(filepath.Join(outputDir, "*.bin"))
	if err != nil {
		panic(err) // Handle error appropriately in real usage
	}

	// Read and concatenate all ABI files
	var mergedABI []byte
	for _, file := range abiFiles {
		abi, err := os.ReadFile(file)
		if err != nil {
			panic(err) // Handle error appropriately in real usage
		}
		mergedABI = append(mergedABI, abi...)
	}

	// Read and concatenate all BIN files
	var mergedBIN []byte
	for _, file := range binFiles {
		bin, err := os.ReadFile(file)
		if err != nil {
			panic(err) // Handle error appropriately in real usage
		}
		mergedBIN = append(mergedBIN, bin...)
	}

	// Write merged ABI and BIN files
	mergedABIFilename := filepath.Join(outputDir, "Merged.abi")
	mergedBINFilename := filepath.Join(outputDir, "Merged.bin")
	err = os.WriteFile(mergedABIFilename, mergedABI, 0600) //nolint:gomnd
	if err != nil {
		panic(err) // Handle error appropriately in real usage
	}
	err = os.WriteFile(mergedBINFilename, mergedBIN, 0600) //nolint:gomnd
	if err != nil {
		panic(err) // Handle error appropriately in real usage
	}
	// Generate Go bindings using abigen
	outFile := filepath.Join(apiDir, generatedFileName+".go")

	cmd = exec.Command("abigen", "--bin="+mergedBINFilename, "--abi="+mergedABIFilename, "--pkg="+generatedPackageName, "--type="+generatedFileName, "--out="+outFile) //nolint:lll,gosec
	output, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to generate Go bindings: %v\nOutput: %s", err, string(output))
	}
	fmt.Println("Go bindings generated successfully")
}
