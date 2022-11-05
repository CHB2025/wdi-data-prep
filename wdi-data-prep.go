package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	csvRegex         = `"(?:\\"|[^"])*"|[^,]*`
	yearStartIndex   = 4
	seriesNameIndex  = 2
	countryNameIndex = 0
	countryCodeIndex = 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a path to the csv file.\n")
		return
	}
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading from %v -- %v", path, err.Error())
	}

	reg := regexp.MustCompile(csvRegex)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	originalColumns := reg.FindAllString(scanner.Text(), -1)
	var data [][]string
	newColumns := []string{
		"Country Name",
		"Country Code",
		"Year",
	}
	for scanner.Scan() {
		row := reg.FindAllString(scanner.Text(), -1)
		data = append(data, row)
		if strings.TrimSpace(row[seriesNameIndex]) != "" && !sliceContains(newColumns, row[seriesNameIndex]) {
			newColumns = append(newColumns, row[seriesNameIndex])
		}
	}
	newRows := make(map[string][]string)

	// Combining every series row for a country-year key into one
	for _, r := range data {
		if len(r) != len(originalColumns) {
			fmt.Printf("Length of row \"%v\" (%v) does not match expected length %v.\n", r, len(r), len(originalColumns))
			break
		}
		if r[seriesNameIndex] == "" {
			break
		}
		for i, yearData := range r[yearStartIndex:] {
			year := strings.Split(originalColumns[yearStartIndex+i], " ")[0]
			key := r[countryCodeIndex] + "_" + year
			if newRows[key] == nil {
				newRows[key] = make([]string, len(newColumns))
				newRows[key][countryNameIndex] = r[0]
				newRows[key][countryCodeIndex] = r[1]
				newRows[key][2] = year
			}
			newRows[key][indexOf(newColumns, r[2])] = yearData
		}
	}

	// Putting the output file together and into a file
	output := fmt.Sprintln(strings.Join(newColumns, ","))
	for _, values := range newRows {
		output += fmt.Sprintln(strings.Join(values, ","))
	}
	outputPath := "output.csv"
	if len(os.Args) >= 3 {
		outputPath = os.Args[2]
	}
	oErr := os.WriteFile(outputPath, []byte(output), 0666)
	if err != nil {
		fmt.Printf("Error writing to %v -- %v", outputPath, oErr.Error())
	}
}

func sliceContains(s []string, v string) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}

func indexOf(s []string, v string) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}
