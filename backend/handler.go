package backend

import (
	"html/template"
	"net/http"
	"strconv"
)

var artists []Artist

func HomePage(w http.ResponseWriter, r *http.Request) { // execute home page html
	switch r.URL.Path {
	case "/":
		var err error
		artists, err = GetArtists()
		if err != nil {
			errPage := ErrorPage{ErrStatus: `Error 500`, ErrMsg: `Internal Server Error`}
			t, _ := template.ParseFiles("frontend/html/errorPage.html")
			w.WriteHeader(500)
			t.Execute(w, errPage)
			return //500
		}
		t, _ := template.ParseFiles("frontend/html/homePage.html")
		t.Execute(w, artists)

	default:
		errPage := ErrorPage{ErrStatus: `Error 404`, ErrMsg: `Page Not Found`}
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(404)
		t.Execute(w, errPage)
		return //404
	}
}

func ArtistsPage(w http.ResponseWriter, r *http.Request) { // get selected artist data from api and present it in artist.html
	artistId := r.FormValue("ArtistID")
	artistIndex, _ := strconv.Atoi(artistId)

	var Band Band
	var err error
	errPage := ErrorPage{ErrStatus: `Error 500`, ErrMsg: `Internal Server Error`}

	if artistIndex-1 <= len(artists) && artistIndex > 0 {
		Band.Artist = artists[artistIndex-1]
	} else {
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(500)
		t.Execute(w, errPage)
		return
	}
	Band.Dates, err = GetDates(artistId)
	if err != nil {
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(500)
		t.Execute(w, errPage)
		return
	}
	Band.Locations, err = GetLocations(artistId)
	if err != nil {
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(500)
		t.Execute(w, errPage)
		return
	}
	Band.RelationsMap, err = GetRelationsMap(artistId)
	if err != nil {
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(500)
		t.Execute(w, errPage)
		return
	}
	Band.Relations = RelationsMapToRelation(Band.RelationsMap) // arrange relations to be used in artist template

	//test coordinates
	Band.Coordinates, err = GetCoordinates(Band.Locations.Locations)
	if err != nil {
		t, _ := template.ParseFiles("frontend/html/errorPage.html")
		w.WriteHeader(500)
		t.Execute(w, errPage)
		return
	}

	//test end

	t, _ := template.ParseFiles("frontend/html/artist.html")
	t.Execute(w, Band)
}
