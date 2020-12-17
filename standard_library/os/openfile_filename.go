package main

import "fmt"
import "log"
import "os"

func main() {
  //Open file named "see", set operator flags "Append to file|CREATE if not exists|WriteRead Only, chmod 0666"
  //Check for err and throw fatal with file name if != nil
	logFile, err := os.OpenFile("./see", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("WE FUCKED UP %s", logFile.Name())
	}
  //Utilize the Name() to return name of OpenFile object
	fmt.Printf("We successfully opened file '%s'\n", logFile.Name())
}

