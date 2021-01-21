package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func createScanner(f *os.File) []string {
	//Create bufio.Scanner obj
	fileScanner := bufio.NewScanner(f)
	//split by new lines
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	//while fileScanner.Scan() returns true, append values as text into slice lines
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func main() {
	//Returns (*os.File) obj, (errors.error) obj
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatalf("Could not obtain *os.File pointer: %+v", err)
	}

	//If we all else fails, run this at the end, no matter what
	defer file.Close()

	//obtain []string from createScanner(), loop through slice
	fileLines := createScanner(file)
	for _, line := range fileLines {
		fmt.Println(line)
	}
}
