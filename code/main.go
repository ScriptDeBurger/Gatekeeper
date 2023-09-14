package main

import (
	"fmt"
	"time"
)

func isTimeBetween(current time.Time, start, end time.Time) bool {
	if start.Before(end) {
		return current.After(start) && current.Before(end)
	}
	return current.After(start) || current.Before(end)
}

func welcomeMessage() string {
	var groet string
	currentTime := time.Now()
	morningBegin := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 7, 0, 0, 0, currentTime.Location())
	morningEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 12, 0, 0, 0, currentTime.Location())
	noonEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 18, 0, 0, 0, currentTime.Location())
	eveningEnd := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())

	if isTimeBetween(currentTime, morningBegin, morningEnd) {
		groet = "Goedemorgen! "
	} else if isTimeBetween(currentTime, morningEnd, noonEnd) {
		groet = "Goedemiddag! "
	} else if isTimeBetween(currentTime, noonEnd, eveningEnd) {
		groet = "Goedenavond! "
	} else {
		groet = "Sorry, de parkeerplaats is 's nachts gesloten."
	}
	return groet
}

func kentekenCheck(kenteken string) string {
	var succes string
	var denied string
	kentekenArray := [5]string{"HK-66-PH", "JP-F3-2D", "CP3-R2-D", "G-43-DS2", "T4-C9-5K"}

	for i := 0; i < len(kentekenArray); i++ {
		if kentekenArray[i] == kenteken {
			succes = "U heeft toegang."
			break
		} else {
			denied = "U heeft helaas geen toegang tot het parkeerterrein."
		}
	}

	if succes != "" {
		return succes
	} else {
		return denied
	}
}

func main() {
	var input string

	fmt.Printf(welcomeMessage())

	fmt.Println("Welkom bij Fonteyn Vakantieparken.")
	fmt.Print("Voer uw kenteken in: ")
	fmt.Scan(&input)
	fmt.Println(kentekenCheck(input))
}
