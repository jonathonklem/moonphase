package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"math"
)

type MoonPhase struct {
	CurrentConditions struct {
		Moonphase float64 `json:"moonphase"`
	} `json:"currentConditions"`
}

func getDaysUntilFullMoon() int {
	apiKey := os.Getenv("API_KEY");
	location := os.Getenv("LOCATION");
	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "?unitGroup=metric&elements=moonphase&contentType=json&key=" + apiKey

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var moonPhase MoonPhase;
	err = json.NewDecoder(response.Body).Decode(&moonPhase)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// api returns 1 = new moon, .5= full moon. 1/28 = 1 day of moonphase
	return int(math.Ceil(((.5-moonPhase.CurrentConditions.Moonphase)/.03571428571)))
}

func sendFullMoonAlert(daysUntilFullMoon int) {
    // Set the email subject and body
    subject := "Full moon is " + fmt.Sprint(daysUntilFullMoon) + " days away!"
    body := "Please be mindful of the upcoming full moon."

	sendEmail(subject, body)
    
}
