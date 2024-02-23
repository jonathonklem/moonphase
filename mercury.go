package main

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)

type RetrogradeResponse	 struct {
	IsRetrograde bool `json:"is_retrograde"`
}

func getMercuryResponseToday() bool {
	retrogradeResponse, err  := http.Get("https://mercuryretrogradeapi.com")
	defer retrogradeResponse.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var retrograde RetrogradeResponse;
	err = json.NewDecoder(retrogradeResponse.Body).Decode(&retrograde)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return retrograde.IsRetrograde
}

func getMercuryResponseNextWeek() bool {
	currentTime := time.Now()
    oneWeekFromNow := currentTime.AddDate(0, 0, 7)
    dateString := oneWeekFromNow.Format("2006-01-02")

	retrogradeResponse, err  := http.Get("https://mercuryretrogradeapi.com?date="+dateString)
	defer retrogradeResponse.Body.Close()

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var retrograde RetrogradeResponse;
	err = json.NewDecoder(retrogradeResponse.Body).Decode(&retrograde)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return retrograde.IsRetrograde
}