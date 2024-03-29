package backend

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []Artist
	response, err := http.Get(url)
	if err != nil {
		// fmt.Println("Error : ", err)
		return artists, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		// fmt.Println("Error : ", err)
		return artists, err
	}
	return artists, err
}

func GetLocations(id string) (Locations, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations/"
	var locations Locations
	response, err := http.Get(url + id)
	if err != nil {
		return locations, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&locations)
	if err != nil {
		return locations, err
	}

	// for i, location := range locations.Locations {
	// 	locations.Locations[i] = regexp.MustCompile("-").ReplaceAllString(location, ", ")
	// }
	return locations, err
}

func GetDates(id string) (Dates, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates/"
	var dates Dates
	response, err := http.Get(url + id)
	if err != nil {
		// fmt.Println("Error : ", err)
		return dates, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&dates)
	if err != nil {
		// fmt.Println("Error : ", err)
		return dates, err
	}
	return dates, err
}

func GetRelationsMap(id string) (RelationsMap, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation/"
	var relationsMap RelationsMap
	response, err := http.Get(url + id)
	if err != nil {
		// fmt.Println("Error : ", err)
		return relationsMap, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&relationsMap)
	if err != nil {
		// fmt.Println("Error : ", err)
		return relationsMap, err
	}
	return relationsMap, err
}

func RelationsMapToRelation(rm RelationsMap) []Relations { // transform relation map from map[string][]string to object Realtions {Location string, Date []string}
	var FinalRelation []Relations
	for location, dates := range rm.DatesLocations {
		newRelation := Relations{Locations: location, Dates: dates}
		FinalRelation = append(FinalRelation, newRelation)
	}
	return FinalRelation
}

func GetCoordinates(locationsNames []string) ([]Coordinates, error) {
	var coordinants []Coordinates
	var thisCoordinates Coordinates
	var res []map[string]interface{}
	var locationsNewNames []string

	for _, location := range locationsNames {
		newName := regexp.MustCompile("-").ReplaceAllString(location, ",")
		locationsNewNames = append(locationsNewNames, newName)
	}

	for _, loc := range locationsNewNames {
		// url := `https://nominatim.openstreetmap.org/search?q=` + loc + `&format=json`
		url := "https://nominatim.openstreetmap.org/search?format=json&q=" + loc
		response, err := http.Get(url)
		if err != nil {
			return coordinants, err
		}
		defer response.Body.Close()

		err = json.NewDecoder(response.Body).Decode(&res)

		if err != nil {
			return coordinants, err
		}

		thisCoordinates.Lat, err = strconv.ParseFloat(res[0]["lat"].(string), 64)
		if err != nil {
			return coordinants, err
		}

		thisCoordinates.Lng, err = strconv.ParseFloat(res[0]["lon"].(string), 64)
		if err != nil {
			return coordinants, err
		}

		coordinants = append(coordinants, thisCoordinates)
	}

	return coordinants, nil
}
