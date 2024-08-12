package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/tidwall/gjson"
)

func getHCLFiles(root string) ([]string, error) {
	var hclFiles []string

	// Walk through the directory tree.
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories or files that contain `.terragrunt-cache` in their path.
		if strings.Contains(path, ".terragrunt-cache") {
			return nil
		}

		// Skip `.terraform.lock.hcl` files.
		if d.Name() == ".terraform.lock.hcl" {
			return nil
		}

		// Check if the file has a .hcl extension.
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".hcl") {
			hclFiles = append(hclFiles, path)
		}
		return nil
	})

	return hclFiles, err
}

// isHCL checks if the file content is valid HCL.
func isHCL(content []byte) bool {
	parser := hclparse.NewParser()
	_, diag := parser.ParseHCL(content, "temp.hcl")
	return !diag.HasErrors()
}

// isJSON checks if the file content is valid JSON.
func isJSON(content []byte) bool {
	return gjson.ValidBytes(content)
}

// hasMisusedColons checks for colon usage in key-value pairs outside of strings or URLs.
func hasMisusedColons(content string) bool {
	// Regex to match colons in key-value pairs that are not inside strings, URLs, or within valid HCL lists/maps.
	re := regexp.MustCompile(`(?m)^\s*[a-zA-Z0-9_-]+\s*:\s*[^"]`)

	// Find all matches in the content.
	matches := re.FindAllString(content, -1)

	// If there are any matches, it indicates potential misuse of colons.
	return len(matches) > 0
}

// checkFileFormat determines the format of the file.
func checkFileFormat(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err)
	}

	content := string(data)

	if isHCL(data) {
		// Custom check: look for misused colons in the HCL file.
		if hasMisusedColons(content) {
			return "HCL with JSON-like syntax"
		}
		return "HCL"
	}

	if isJSON(data) {
		return "JSON"
	}

	return "Unknown format"
}

func main() {
	// Define the directory where the terragrunt files are located.
	dir := "infra"
	hclFiles, err := getHCLFiles(dir)
	if err != nil {
		log.Fatalf("Failed to get files with `.hcl` extension: %v", err)
	}

	for _, file := range hclFiles {
		format := checkFileFormat(file)
		if format != "HCL" {
			fmt.Printf("File: %s is in %s format.\n", file, format)
		}
	}
}
