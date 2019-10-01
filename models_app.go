package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

type apps struct {
	AppArray []app `json:"apps"`
}

type app struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	URLPlay     string `json:"urlPlay"`
	URLAmazon   string `json:"urlAmazon"`
	URLApk      string `json:"urlApk"`
	SortOrder   int    `json:"sortOrder"`
}

var appData apps

func init() {
	appData = loadApps()
	sort.Slice(appData.AppArray, func(i, j int) bool {
		return appData.AppArray[i].SortOrder < appData.AppArray[j].SortOrder
	})
	for i := 0; i < len(appData.AppArray); i++ {
		log.Printf("App: " + strconv.Itoa(appData.AppArray[i].SortOrder) + " | " + strconv.Itoa(appData.AppArray[i].ID) + " | " + appData.AppArray[i].Name)
	}
}

func loadApps() apps {
	appsJSON, err := os.Open("data/apps.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer appsJSON.Close()
	var apps apps
	byteValue, _ := ioutil.ReadAll(appsJSON)
	json.Unmarshal(byteValue, &apps)
	return apps
}

func getAllApps() []app {
	return appData.AppArray
}

func getAppByID(id int) (*app, error) {
	for _, a := range appData.AppArray {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("App not found")
}
