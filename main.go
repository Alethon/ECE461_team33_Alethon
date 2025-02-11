package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	//https://linguinecode.com/post/how-to-import-local-files-packages-in-golang
	"github.com/KevinMi2023p/ECE461_TEAM33/package_analyzer"
)

// main function will handle the command line arguments
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(1)
	}

	var argument string = os.Args[1]
	// Variable that'll eventually be used as the final variable
	final_output := []package_analyzer.Metrics{}

	// Section uses the file arugment
	file, err := os.Open(argument)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	// Setting up scanner to read line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Analyze each line
	for scanner.Scan() {
		url_on_line := scanner.Text()
		datapoint := package_analyzer.Analyze(url_on_line)
		// function will return "nil" if it isn't a valid URL
		if datapoint != nil {
			final_output = append(final_output, *datapoint)
		}
	}

	// Sort what we're about to output
	sort.Slice(final_output, func(i, j int) bool {
		return final_output[i].Net_score >= final_output[j].Net_score
	})

	// Prints out the output, with each datapoint having it's own line
	for _, obj := range final_output {
		// fmt.Println(package_analyzer.Metrics_toString(obj))
		fmt.Println(package_analyzer.Metrics_toString(obj))
	}

	os.Exit(0)
}
