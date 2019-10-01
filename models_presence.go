package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

type presences struct {
	PresenceArray []presence `json:"presences"`
}

type presence struct {
	Label     string `json:"label"`
	URL       string `json:"url"`
	SortOrder int    `json:"sortOrder"`
}

var presenceData presences

func init() {
	presenceData = loadPresences()
	sort.Slice(presenceData.PresenceArray, func(i, j int) bool {
		return presenceData.PresenceArray[i].SortOrder < presenceData.PresenceArray[j].SortOrder
	})
	for i := 0; i < len(presenceData.PresenceArray); i++ {
		log.Printf("Presence: " + strconv.Itoa(presenceData.PresenceArray[i].SortOrder) + " | " + presenceData.PresenceArray[i].Label + " | " + presenceData.PresenceArray[i].URL)
	}
}

func loadPresences() presences {
	presencesJSON, err := os.Open("data/presences.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer presencesJSON.Close()
	var presences presences
	byteValue, _ := ioutil.ReadAll(presencesJSON)
	json.Unmarshal(byteValue, &presences)
	return presences
}

func getAllPresences() []presence {
	return presenceData.PresenceArray
}
