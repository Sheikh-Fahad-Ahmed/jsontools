package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/divakarans/setjson/setter"
	"github.com/Sheikh-Fahad-Ahmed/jsontools/internal/helper"
	"github.com/tidwall/gjson"
	"github.com/vharshitha0089/json2yaml/converter"
)

func main() {
	// ------------ Command-line arguments ------------
	inputPath := flag.String("input", "default", "provide an input file (required)")
	outputPath := flag.String("output", "out.yaml", "provide an output file")
	flag.Parse()

	if *inputPath == "--output" {
		fmt.Println("jsontool --input <file> [--output <file>]")
		os.Exit(1)
	}


	// ------------ Load ------------

	jsonByte, err := os.ReadFile(*inputPath)
	helper.CheckErr(err)

	jsonStr := string(jsonByte)

	fieldCount := gjson.Parse(jsonStr).Get("@keys.#")


	fmt.Printf("\n\nLoaded: %s (%d fields)\n\n", *inputPath, fieldCount.Int())

	// ------------ Verify Step ------------
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("── Verify field ──")
	fieldName := helper.Prompt(scanner, "Field to verify:")
	result := gjson.Get(jsonStr, fieldName)
	if !result.Exists() {
		fmt.Printf("Field %s not found\n", fieldName)
	} else {
		fmt.Printf("Current value %s\n", result.String())
		answer := helper.Prompt(scanner, "Is this correct? (y/n): ")
		if answer == "n" {
			newVal := helper.Prompt(scanner, "Enter new value:")
			jsonStr = setter.Setjson(jsonStr, fieldName, newVal) //------------------------------------
			fmt.Printf("✓ Updated %s → %s\n\n", fieldName, newVal)
		}
	}

	// ------------ Update Field ------------
	fmt.Println("── Update field ──")
	fieldName = helper.Prompt(scanner, "Field to update: ")
	result = gjson.Get(jsonStr, fieldName)
	if !result.Exists() {
		fmt.Printf("Field %s not found", fieldName)
	} else {
		fmt.Printf("Current value: %s\n", result.String())
		newVal := helper.Prompt(scanner, "Enter new value: ")
		jsonStr = setter.Setjson(jsonStr, fieldName, newVal) //------------------------------------
		fmt.Printf("✓ Updated %s → %s\n\n", fieldName, newVal)
	}

	// ------------ Convert To YAML ------------
	fmt.Println("Converting to YAML...")
	yamlStr := converter.Convert(jsonStr)

	err = os.WriteFile(*outputPath, []byte(yamlStr), 0644)
	helper.CheckErr(err)
	fmt.Printf("✓ Wrote %s\n", *outputPath)
	fmt.Printf("\nDone.\n\n")
}

